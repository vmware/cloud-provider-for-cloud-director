package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/vmware/go-vcloud-director/v2/govcd"
	"github.com/vmware/go-vcloud-director/v2/util"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	logFile = "token_verification.log"
)

const (
	VCloudApiVersion = "36.0"
)

func main() {
	var err error

	defaultCmd := flag.NewFlagSet("default", flag.ExitOnError)
	verifyCmd := flag.NewFlagSet("verify-token", flag.ExitOnError)

	var vcdHost string
	defaultCmd.StringVar(&vcdHost, "vcd-host", "", "VCD Site URL (Mandatory)")
	verifyCmd.StringVar(&vcdHost, "vcd-host", "", "VCD Site URL (Mandatory)")

	var orgName string
	defaultCmd.StringVar(&orgName, "org", "system", "orgName to which the user belongs to (use 'system' or system org)")
	verifyCmd.StringVar(&orgName, "org", "system", "orgName to which the user belongs to (use 'system' or system org)")

	var username string
	defaultCmd.StringVar(&username, "user", "username", "username of the user account being used")
	verifyCmd.StringVar(&username, "user", "username", "username of the user account being used")

	var password string
	defaultCmd.StringVar(&password, "password", "password", "password of the user account being used")
	verifyCmd.StringVar(&password, "password", "password", "password of the user account being used")

	var oauthClientName string
	defaultCmd.StringVar(&oauthClientName, "client-name", "oauthclient", "oauth client name to create the refresh token")
	verifyCmd.StringVar(&oauthClientName, "client-name", "oauthclient", "oauth client name to create the refresh token")

	var insecure bool
	defaultCmd.BoolVar(&insecure, "skip-verify", false, "skip verify ")
	verifyCmd.BoolVar(&insecure, "skip-verify", false, "skip verify ")

	var refreshToken string
	defaultCmd.StringVar(&refreshToken, "refresh-token", "", "refresh token value - used when running with verify-token sub-command")
	verifyCmd.StringVar(&refreshToken, "refresh-token", "", "refresh token value - used when running with verify-token sub-command")

	flag.Parse()

	util.EnableLogging = true
	util.ApiLogFileName = logFile
	util.SetLog()
	util.LogHttpRequest = true
	util.LogHttpResponse = true
	util.LogPasswords = false

	switch os.Args[1] {
	case "default":
		defaultCmd.Parse(os.Args[2:])
		if vcdHost == "" {
			flag.Usage()
			panic(fmt.Errorf("--vcd-host is a mandatory parameter"))
		}

		if orgName == "" {
			flag.Usage()
			panic(fmt.Errorf("--org is a mandatory parameter"))
		}

		fmt.Println("creating refresh token ...")
		// create refresh token
		refreshToken, clientID, err := createRefreshToken(vcdHost, orgName, username, password, oauthClientName, insecure)
		if err != nil {
			panic(fmt.Errorf("failed to create refresh token: [%v]", err))
		}

		// verify refresh token
		fmt.Println("verifying refresh token ...")
		err = verifyRefreshToken(vcdHost, orgName, refreshToken, insecure)
		if err != nil {
			panic(fmt.Errorf("failed to verify refresh token: [%v]", err))
		}

		if username == "" {
			flag.Usage()
			panic(fmt.Errorf("--user is a mandatory parameter"))
		}

		if oauthClientName == "" {
			flag.Usage()
			panic(fmt.Errorf("--client-name is a mandatory parameter"))
		}

		fmt.Println("deleting the refresh token ...")
		err = deleteRefreshToken(vcdHost, orgName, username, password, clientID, insecure)

		fmt.Println("successfully authenticated through refresh token")
	case "verify-token":
		verifyCmd.Parse(os.Args[2:])
		if vcdHost == "" {
			flag.Usage()
			panic(fmt.Errorf("--vcd-host is a mandatory parameter"))
		}

		if orgName == "" {
			flag.Usage()
			panic(fmt.Errorf("--org is a mandatory parameter"))
		}
		if refreshToken == "" {
			flag.Usage()
			panic(fmt.Errorf("please specify the refresh token using --refresh-token option"))
		}

		// verify refresh token
		fmt.Println("verifying refresh token ...")
		err = verifyRefreshToken(vcdHost, orgName, refreshToken, insecure)
		if err != nil {
			panic(fmt.Errorf("failed to verify refresh token: [%v]", err))
		}
	default:
		panic(fmt.Errorf("invalid sub-command"))
	}
}

func createRefreshToken(vcdHost, orgName, username, password, oauthClientName string, insecure bool) (string, string, error) {
	href := fmt.Sprintf("%s/api", vcdHost)
	u, err := url.ParseRequestURI(href)
	if err != nil {
		return "", "", fmt.Errorf("unable to parse url [%s]: %s", href, err)
	}
	vcdClient := govcd.NewVCDClient(*u, insecure)
	vcdClient.Client.APIVersion = VCloudApiVersion

	resp, err := vcdClient.GetAuthResponse(username, password, orgName)
	if err != nil {
		return "", "", fmt.Errorf("unable to authenticate [%s/%s] for url [%s]: [%+v] : [%v]",
			orgName, username, href, resp, err)
	}

	// create client request url
	createOAuthClientRequestHref := fmt.Sprintf("%s/oauth/tenant/%s/register", vcdHost, orgName)
	if strings.ToLower(orgName) == "system" {
		createOAuthClientRequestHref = fmt.Sprintf("%s/oauth/provider/register", vcdHost)
	}
	createOauthClientUrl, err := url.ParseRequestURI(createOAuthClientRequestHref)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse url to create oauth client: [%v]", err)
	}
	oauthClientRequestBody := map[string]interface{}{
		"client_name": oauthClientName,
	}
	b := new(bytes.Buffer)
	if err = json.NewEncoder(b).Encode(oauthClientRequestBody); err != nil {
		return "", "", fmt.Errorf("error creating request body to create oauth client [%s]", oauthClientName)
	}
	vcdClient.Client.SetCustomHeader(map[string]string{
		"Accept":       "application/*;version=36.0",
		"Content-Type": "application/json",
	})
	createOauthReq := vcdClient.Client.NewRequest(nil, http.MethodPost, *createOauthClientUrl, b)
	createOauthResp, err := vcdClient.Client.Http.Do(createOauthReq)
	if err != nil {
		return "", "", fmt.Errorf("error creating oauth client [%s]: [%v]", oauthClientName, err)
	}
	if createOauthResp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("invalid response creating oauth client [%s]: [%v]", oauthClientName, err)
	}
	var createOauthRespBody map[string]interface{}
	err = json.NewDecoder(createOauthResp.Body).Decode(&createOauthRespBody)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse response body to map[string]interface{}: [%v]", err)
	}
	clientID := createOauthRespBody["client_id"].(string)

	// create refresh token
	createApiTokenHref := fmt.Sprintf("%s/oauth/tenant/%s/token", vcdHost, orgName)
	if strings.ToLower(orgName) == "system" {
		createApiTokenHref = fmt.Sprintf("%s/oauth/provider/token", vcdHost)
	}
	createRefreshTokenUrl, err := url.ParseRequestURI(createApiTokenHref)
	data := bytes.NewBufferString(fmt.Sprintf("client_id=%s&grant_type=urn:ietf:params:oauth:grant-type:jwt-bearer&assertion=%s", clientID, vcdClient.Client.VCDToken))

	vcdClient.Client.RemoveCustomHeader()
	vcdClient.Client.SetCustomHeader(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"Accept":       "application/*;version=36.0",
	})
	createApiTokenReq := vcdClient.Client.NewRequest(nil, http.MethodPost, *createRefreshTokenUrl, data)
	createApiTokenResp, err := vcdClient.Client.Http.Do(createApiTokenReq)
	if err != nil {
		return "", "", fmt.Errorf("error creating a refresh token: [%v]", err)
	}
	if createApiTokenResp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("invalid response creating a refresh token: [%d]", createApiTokenResp.StatusCode)
	}
	var apiTokenRespBody map[string]interface{}
	err = json.NewDecoder(createApiTokenResp.Body).Decode(&apiTokenRespBody)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse response body to map[string]interface{}: [%v]", err)
	}
	vcdClient.Client.RemoveCustomHeader()
	fmt.Println("Created refersh token with client name", oauthClientName, " and expiring on ", apiTokenRespBody["expires_in"].(float64))
	return apiTokenRespBody["refresh_token"].(string), clientID, nil
}

func verifyRefreshToken(vcdHost, orgName, refreshToken string, insecure bool) error {
	href := fmt.Sprintf("%s/api", vcdHost)
	u, err := url.ParseRequestURI(href)
	if err != nil {
		return fmt.Errorf("unable to parse url [%s]: %s", href, err)
	}
	vcdClient := govcd.NewVCDClient(*u, insecure)
	vcdClient.Client.APIVersion = VCloudApiVersion

	if err = vcdClient.SetToken(orgName, govcd.ApiTokenHeader, refreshToken); err != nil {
		return fmt.Errorf("failed to set api token to vcd client: [%v]", err)
	}
	return nil
}

func deleteRefreshToken(vcdHost, orgName, username, password, clientID string, insecure bool) error {
	href := fmt.Sprintf("%s/api", vcdHost)
	u, err := url.ParseRequestURI(href)
	if err != nil {
		return fmt.Errorf("unable to parse url [%s]: %s", href, err)
	}
	vcdClient := govcd.NewVCDClient(*u, insecure)
	vcdClient.Client.APIVersion = VCloudApiVersion

	resp, err := vcdClient.GetAuthResponse(username, password, orgName)
	if err != nil {
		return fmt.Errorf("unable to authenticate [%s/%s] for url [%s]: [%+v] : [%v]",
			orgName, username, href, resp, err)
	}

	currentSessionUrl, err := vcdClient.Client.OpenApiBuildEndpoint(fmt.Sprintf("1.0.0/tokens/urn:vcloud:token:%s", clientID))
	if err != nil {
		return fmt.Errorf("failed to construct delete token url [%v]", err)
	}

	err = vcdClient.Client.OpenApiDeleteItem(vcdClient.Client.APIVersion, currentSessionUrl, url.Values{}, nil)
	if err != nil {
		return fmt.Errorf("error while deleting refresh token [%v]", err)
	}
	return nil
}
