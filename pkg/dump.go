package pkg

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	DUMP_OUTPUT_PROPERTIES = "properties"
)

type versionContext string

var (
	ReleasedVersionContext    = "released"
	IterationVersionContext   = "iteration"
	IntegrationVersionContext = "integration"
	PullRequestVersionContext = "pullrequest"
)

type dumpVersionData struct {
	Context string
	Branch  string
	Tag     string

	// SourceBranch string
	// TargetBranch string

	CommitID      string
	CommitShortID string

	Timestamp string

	IterationVersion string
	ReleasedVersion  string
	IncreasingNum    int

	PRID int64

	GitDescribeVersion string
}

func (d dumpVersionData) String() string {
	bf := bytes.NewBufferString("")
	bf.WriteString("VERSION_CONTEXT=" + d.Context + "\n")
	bf.WriteString("BRANCH=" + d.Branch + "\n")
	bf.WriteString("TAG=" + d.Tag + "\n")
	bf.WriteString("COMMIT_ID=" + d.CommitID + "\n")
	bf.WriteString("COMMIT_SHORT_ID=" + d.CommitShortID + "\n")
	bf.WriteString("TIMESTAMP=" + d.Timestamp + "\n")

	bf.WriteString("PR_ID=" + fmt.Sprintf("%d", d.PRID) + "\n")
	bf.WriteString("GIT_DESCRIBE_VERSION=" + d.GitDescribeVersion + "\n")
	// increasing number is depencency on all tags
	// bf.WriteString("INCREASING_NUM=" + fmt.Sprintf("%d", d.IncreasingNum) + "\n")
	// depend on all branches
	// bf.WriteString("ITERATION_VERSION=" + d.IterationVersion + "\n")
	// bf.WriteString("RELEASED_VERSION=" + d.ReleasedVersion + "\n")

	return bf.String()
}

func Dump() (res string, err error) {
	result := dumpVersionData{
		Timestamp: time.Now().Format("01021504"),
	}

	err = fetchInfo(&result)
	if err != nil {
		return "", err
	}

	branches, err := GetAllBranches()
	if err != nil {
		return "", err
	}

	releasedPattern, _ := regexp.Compile(`release-\d+\.\d+`)
	var maxV float64 = 1.0
	for _, item := range branches {
		matched := releasedPattern.MatchString(item)
		if matched {
			v, err := strconv.ParseFloat(strings.TrimPrefix(item, "release-"), 32)
			if err != nil {
				fmt.Printf("parse to int error: %s, skip\n", err.Error())
				continue
			}

			if v > maxV {
				maxV = v
			}
		}
	}

	result.IterationVersion = fmt.Sprintf("%.1f", maxV+0.1)
	if maxV != 1.0 {
		result.ReleasedVersion = fmt.Sprintf("%.1f", maxV) // TODO : get target branch
	}

	result.GitDescribeVersion, err = GitDescribe("--dirty", "--always", "--tags", "--long", "--match", fmt.Sprintf("v%f", maxV))
	if err != nil {
		return "", err
	}

	increasingNum := 0
	tags, err := GetAllTags()
	if err != nil {
		return "", err
	}
	for _, item := range tags {
		if strings.HasPrefix(item, "v"+result.ReleasedVersion) {
			patchVersionStr := strings.TrimPrefix(item, "v"+result.ReleasedVersion+".")
			patchVersion, err := strconv.ParseInt(patchVersionStr, 10, 8)
			if err != nil {
				fmt.Printf("parse to int error: %s, skip\n", err.Error())
				continue
			}
			if patchVersion > int64(increasingNum) {
				increasingNum = int(patchVersion)
			}
		}
	}

	result.IncreasingNum = increasingNum + 1

	return result.String(), nil
}

func fetchInfo(result *dumpVersionData) (err error) {
	bts, err := os.ReadFile(".git/FETCH_HEAD")
	if err != nil {
		return err
	}

	fetchHead := string(bts)
	fetchHead = strings.ReplaceAll(fetchHead, "\t", " ")

	var segments = []string{}
	for _, item := range strings.Split(fetchHead, " ") {
		if strings.TrimSpace(item) != "" {
			segments = append(segments, strings.TrimSpace(item))
		}
	}
	if len(segments) < 2 {
		return fmt.Errorf("error format of fetch head: %s", fetchHead)
	}
	result.CommitID = segments[0]
	result.CommitShortID = result.CommitID[0:8]

	refs := segments[1]
	refs = strings.TrimPrefix(refs, "'")
	refs = strings.TrimSuffix(refs, "'")

	if strings.HasPrefix(refs, "refs/merge-requests/") || strings.HasPrefix(refs, "refs/pulls/") {
		result.Context = PullRequestVersionContext
		prIDStr := strings.Split(refs, "/")[2]
		result.PRID, err = strconv.ParseInt(prIDStr, 10, 8)
		if err != nil {
			return err
		}
	}

	if strings.HasPrefix(refs, "refs/heads") {
		result.Branch = strings.Split(refs, "/")[2]
	}

	if strings.HasPrefix(refs, "refs/tags") {
		result.Tag = strings.Split(refs, "/")[2]
	}

	return
}
