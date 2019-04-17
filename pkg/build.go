package pkg

import (
	"fmt"
	"strings"
)

// BuildFunc build generating function
type BuildFunc func([]string) (string, error)

func BuildInplace(args []string) (version string, err error) {
	if len(args) != 2 {
		err = fmt.Errorf(`please provide a minor version number to generate build number and the current version. e.g: v0.1 v0.2.0`)
	}
	desiredVersion := args[0]
	tags := []string{args[1]}
	version = GetBuilderVersion(tags, desiredVersion)
	return
}

func BuildGit(args []string) (version string, err error) {
	if len(args) != 1 {
		err = fmt.Errorf(`please provide a minor version number to generate build number. e.g: v0.1`)
		return
	}
	desiredVersion := args[0]
	var tags []string
	tags, err = GetAllTags()
	if err != nil {
		return
	}
	version = GetBuilderVersion(tags, desiredVersion)
	return
}

func GetBuilderVersion(tags []string, version string) (result string) {
	result = fmt.Sprintf("%v-b.1", version)
	if len(tags) == 0 || (len(tags) == 1 && tags[0] == version) {
		return
	}
	prefix := fmt.Sprintf("%v-b.", version)
	tags = FilterTags(prefix, tags, strings.HasPrefix)
	if len(tags) == 0 || (len(tags) == 1 && tags[0] == version) {
		return
	}
	highest := GetHighestPatch(tags, func(build string) string {
		return strings.Replace(build, "-b", "", -1)
	})
	highest++
	result = fmt.Sprintf("%v-b.%d", version, highest)
	return
}
