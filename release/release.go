package release

import (
	_ "embed"
)

const (
	CloudControllerManagerName = "cloud-controller-manager"
)

//go:embed version
var CpiVersion string

