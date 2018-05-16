package pkg

import (
	"strings"

	"github.com/blang/semver"
)

func GetNextChartVersion(current, next string) (ver string, err error) {
	current = strings.TrimPrefix(current, "v")
	next = strings.TrimPrefix(next, "v")
	// increseOnCurrent := false
	// if they are the same should keep the same
	defer func() {
		if ver != "" {
			ver = "v" + ver
		}
	}()
	if current == next {
		ver = current
		return
	}
	if next == "" {
		// increseOnCurrent = true
		next = current
	}
	var (
		curVer, nextVer semver.Version
		// currentMinor, nextMinor string
		// split                   []string
	)
	if curVer, err = semver.Make(current); err != nil {
		return
	}
	switch strings.Count(next, ".") {
	case 1:
		next = next + ".0"
	case 2:
		// perfect
	}
	if nextVer, err = semver.Make(next); err != nil {
		return
	}
	if nextVer.GT(curVer) {
		ver = nextVer.String()
		return
	}
	curVer.Patch++
	ver = curVer.String()
	// if nextVer.EQ(curVer) {

	// 	return
	// }

	return
}
