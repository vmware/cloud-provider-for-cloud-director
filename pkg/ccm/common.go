package ccm

import (
	"fmt"
	"strings"
)

func getUUIDFromProviderID(providerID string) string {
	withoutPrefix := strings.TrimPrefix(providerID, ProviderName+"://")
	return strings.ToLower(strings.TrimSpace(withoutPrefix))
}

func getProviderIDFromUUID(vmUUID string) string {
	return fmt.Sprintf("%s://%s", ProviderName, vmUUID)
}
