package pkg

import (
	"fmt"
	"strings"
)

// BuildFunc build generating function
type BuildFunc func([]string, func() ([]string, error)) (string, error)

func BuildInplace(args []string, _ func() ([]string, error)) (version string, err error) {
	if len(args) != 2 {
		err = fmt.Errorf(`please provide a minor version number to generate build number and the current version. e.g: v0.1 v0.2.0`)
	}
	desiredVersion := args[0]
	tags := []string{args[1]}
	version = GetBuilderVersion(tags, desiredVersion)
	return
}

func BuildGit(args []string, getTagsFunc func() ([]string, error)) (version string, err error) {
	if len(args) < 1 || len(args) > 2 {
		err = fmt.Errorf(`please provide a minor version number to generate build number. e.g: v0.1`)
		return
	}
	desiredVersion := args[0]
	tagPrefix := ""
	if len(args) == 2 {
		tagPrefix = args[1]
	}
	var tags []string
	tags, err = getTagsFunc()
	if err != nil {
		return
	}
	prefixedTags := make([]string, 0, len(tags))
	// filter tags for prefix
	if tagPrefix != "" {
		if !strings.HasSuffix(tagPrefix, "-") {
			tagPrefix += "-"
		}
		for i, tag := range tags {
			if strings.HasPrefix(tags[i], tagPrefix) {
				prefixedTags = append(prefixedTags, strings.TrimPrefix(tag, tagPrefix))
			}
		}
	} else {
		prefixedTags = tags
	}
	version = GetBuilderVersion(prefixedTags, desiredVersion)
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
