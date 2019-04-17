package pkg

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func GenReleaseCandidate(args []string) (version string, err error) {
	switch len(args) {
	case 2:
		version, err = GenRC(args[0], args[1])
	default:
		err = fmt.Errorf("Number of arguments is invalid, should be current version, expected version prefix. e.g v0.1.2 v0.2")
	}
	return
}

var rcRegex = regexp.MustCompile(`v(\d+)\.(\d+|\d+\-[a-zA-Z]+)\.(\d+)`)
var majorMinorRegex = regexp.MustCompile(`v(\d+)\.(\d+)`)

func GenRC(current, desired string) (result string, err error) {
	//
	if !rcRegex.MatchString(current) {
		err = fmt.Errorf("Current \"%s\" should be either a specific version (v0.1.2) or an rc version (v0.1-rc.0)", current)
		return
	}
	if !majorMinorRegex.MatchString(desired) {
		err = fmt.Errorf("Desired \"%s\" should be a vMajor.Minor (v0.1)", desired)
		return
	}
	currentMatches := rcRegex.FindStringSubmatch(current)
	desiredMatches := majorMinorRegex.FindStringSubmatch(desired)
	// if minor and major versions are equal and current version is rc then
	// we should use current version to upgrade
	if currentMatches[1] == desiredMatches[1] &&
		strings.Replace(currentMatches[2], "-rc", "", -1) == desiredMatches[2] &&
		strings.Contains(currentMatches[2], "-rc") {
		var number int
		number, err = strconv.Atoi(currentMatches[3])
		if err != nil {
			return
		}
		number++
		currentMatches[3] = strconv.Itoa(number)
		new := currentMatches[1:]
		result = fmt.Sprintf("v%s", strings.Join(new, "."))
	} else {
		// otherwise we just ignore the current version and upgrade the desired version
		result = fmt.Sprintf("%s-rc.0", desired)
	}
	return
}
