# token-log-collector

## Generate linux binary

To generate the token-log-collector linux binary, execute the following command

```shell
make generate-binary
```

The above make target stores the binary in `./bin` directory. This binary can be used in a linux environment

## Using the binary to collect debug API token related failure

Use the token-log-collector binary to create a sample API token and verify authentication to VCD
```shell
token-log-collector default --vcd-host https://<VCD_SITE> -org <ORG> -user <USER_ORG> -password <PASSWORD> --skip-verify --client-name cseclient
```

Use the token-log-collector binary to verify if an API token is valid
```shell
token-log-collector verify-token --vcd-host https://<VCD_SITE> -org cluster-org -user user  --skip-verify --refresh-token <API_TOKEN>
```