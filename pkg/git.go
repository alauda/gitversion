package pkg

import (
	"fmt"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

// GetAllTags list all tags in the current directory
func GetAllTags() (tags []string, err error) {
	var output []byte
	if output, err = exec.Command("ls", ".git/refs/tags").CombinedOutput(); err != nil {
		return
	}
	text := string(output)
	tags = strings.Split(text, "\n")
	for i := 0; i < len(tags); i++ {
		if tags[i] == "" {
			tags = append(tags[:i], tags[i+1:]...)
		}
	}
	return
}

// GetAllBranches list all branches in the current directory
func GetAllBranches() (branches []string, err error) {
	var output []byte
	if output, err = exec.Command("ls", ".git/refs/heads").CombinedOutput(); err != nil {
		return
	}
	text := string(output)
	branches = strings.Split(text, "\n")
	for i := 0; i < len(branches); i++ {
		if branches[i] == "" {
			branches = append(branches[:i], branches[i+1:]...)
		}
	}
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

// GetHighestPatch return the highest patch number from all tags
func GetHighestPatch(tags []string, transform ...func(string) string) int {
	numbers := make([]int, 0, len(tags))
	for _, t := range tags {
		for _, trans := range transform {
			t = trans(t)
		}
		spl := strings.Split(t, ".")

		num, err := strconv.Atoi(spl[len(spl)-1])
		if err == nil {
			numbers = append(numbers, num)
		}
	}
	if len(numbers) > 0 {
		sort.Ints(numbers)
		return numbers[len(numbers)-1]
	}
	return -1
}

// GitDescribe describe a git using given arguments
func GitDescribe(arguments ...string) (version string, err error) {
	var data []byte
	data, err = exec.Command("git",
		append(
			[]string{"describe"},
			arguments...,
		)...,
	).CombinedOutput()
	if err != nil {
		fmt.Println("error", string(data))
		return
	}
	version = strings.ReplaceAll(string(data), "\n", "")
	return
}
