package utils

import "strings"

// TrimTrailingSlash is used to trimming the final slash of directory/path's passed in.
// Such that it can support paths like /tmp/ and /tmp as the inputs, and create a file output as /tmp/file.txt
func TrimTrailingSlash(path string) string {
	if strings.HasSuffix(path, "/") {
		return path[:len(path)-1]
	}
	return path
}
