package pkg

import (
	"os/exec"
	"strings"
)

// GetAllTags list all tags in the current directory
func GetAllTags() (tags []string, err error) {
	var output []byte
	if output, err = exec.Command("git", "tags").Output(); err != nil {
		return
	}
	text := string(output)
	tags = strings.Split(text, "\n")
	return
}

// FilterTags filter a tag list based on a condition and a filter function
func FilterTags(condition string, tags []string, filter func(string, string) bool) (res []string) {
	if len(tags) == 0 || len(condition) == 0 {
		return tags
	}
	res = make([]string, 0, len(condition))
	for _, r := range tags {
		if filter(r, condition) {
			res = append(res, r)
		}
	}
	return
}
