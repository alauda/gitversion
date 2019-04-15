package pkg

import (
	"fmt"
	"strings"
)

// BumpMinor increses minor version by one
func BumpMinor(minor string) (string, error) {
	version := strings.Replace(minor, "v", "", -1)
	split := strings.Split(version, ".")
	if len(split) != 2 {
		return "", fmt.Errorf("minor version \"%s\" format is incorrect", minor)
	}
	highest := GetHighestPatch([]string{version})
	highest++

	return fmt.Sprintf("v%s.%d", split[0], highest), nil
}
