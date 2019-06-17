package pkg_test

import (
	"testing"

	"github.com/alauda/gitversion/pkg"
)

func TestBuildNumbers(t *testing.T) {

	type TestCase struct {
		Tags     []string
		Version  string
		Prefix   string
		Expected string
	}

	table := []TestCase{
		{
			[]string{"v0.1", "v0.1.1"},
			"v0.1",
			"",
			"v0.1-b.1",
		},
		{
			[]string{"v0.1", "v0.1.1", "v0.1-b.1"},
			"v0.1",
			"",
			"v0.1-b.2",
		},
		{
			[]string{"v0.1", "v0.1.1", "v0.1-b.1"},
			"v0.2",
			"",
			"v0.2-b.1",
		},
		{
			[]string{},
			"v0.2",
			"",
			"v0.2-b.1",
		},
		{
			[]string{"v1.2-b.1"},
			"v1.2",
			"",
			"v1.2-b.2",
		},
	}

	for i, test := range table {
		result := pkg.GetBuilderVersion(test.Tags, test.Version)
		if result != test.Expected {
			t.Errorf("Test %d failed. \"%s\" != \"%s\" ", i, result, test.Expected)
		}
	}
}

func TestBuildGit(t *testing.T) {

	type TestCase struct {
		Tags     []string
		Args     []string
		Expected string
	}

	table := []TestCase{
		{ // without prefix
			[]string{"some", "stable-aaa", "v0.1-b.1", "alauda-devops-v0.1.1"},
			[]string{"v0.1"},
			"v0.1-b.2",
		},
		{ // adding - to the tag prefix
			[]string{"some", "stable-aaa", "alauda-devops-v0.1", "alauda-devops-v0.1.1"},
			[]string{"v0.1", "alauda-devops-"},
			"v0.1-b.1",
		},
		{ // without adding - to the prefix
			[]string{"alauda-devops-v0.1", "alauda-devops-v0.1.1", "alauda-devops-v0.1-b.1"},
			[]string{"v0.1", "alauda-devops"},
			"v0.1-b.2",
		},
	}

	for i, test := range table {
		result, err := pkg.BuildGit(test.Args, func() ([]string, error) {
			return test.Tags, nil
		})
		if err != nil {
			t.Errorf("Test %d failed. should not return error: %v", i, err)
			continue
		}
		if result != test.Expected {
			t.Errorf("Test %d failed. \"%s\" != \"%s\" ", i, result, test.Expected)
		}
	}
}
