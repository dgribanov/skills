package version

import (
	"fmt"
	"regexp"
	"strconv"
)

// GoMinorVersion ищет минорную версию go в выводе "go version" формата "go version go1.19.2 linux/amd64"
func GoMinorVersion(versionOutput []byte) (int, error) {
	versionRegexp := regexp.MustCompile(`go\d+\.(\d+)`)
	submatches := versionRegexp.FindSubmatch(versionOutput)
	if len(submatches) != 2 {
		return 0, fmt.Errorf("не могу найти минорную версию: %s", string(versionOutput))
	}

	minorVersion := submatches[1]
	return strconv.Atoi(string(minorVersion))
}