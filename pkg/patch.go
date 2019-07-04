package pkg

import (
	"errors"
	"fmt"
	"strings"
)

// PatchVersion patch a version
func PatchVersion(args, tags []string) (version string, err error) {

	if len(args) != 1 {
		err = errors.New(`please provide a minor version number to generate patch number. e.g: v0.1`)
		return
	}
	version = args[0]
	if tags == nil {
		tags, err = GetAllTags()
		if err != nil {
			return
		}
	}
	if strings.HasSuffix(version, ".") {
		version = strings.TrimSuffix(version, ".")
	}
	filter := version + "."
	tags = FilterTags(filter, tags, strings.HasPrefix)
	if len(tags) == 0 || (len(tags) == 1 && tags[0] == version) {
		version = fmt.Sprintf("%v.0", version)
		return
	}
	highest := GetHighestPatch(tags)
	highest++
	version = fmt.Sprintf("%v.%d", version, highest)
	return
}
