package cmd_test

import (
	"testing"

	"github.com/alauda/gitversion/cmd"
)

func TestBuildNumbers(t *testing.T) {

	type TestCase struct {
		Tags     []string
		Version  string
		Expected string
	}

	table := []TestCase{
		{
			[]string{"v0.1", "v0.1.1"},
			"v0.1",
			"v0.1.b-1",
		},
		{
			[]string{"v0.1", "v0.1.1", "v0.1.b-1"},
			"v0.1",
			"v0.1.b-2",
		},
		{
			[]string{"v0.1", "v0.1.1", "v0.1.b-1"},
			"v0.2",
			"v0.2.b-1",
		},
		{
			[]string{},
			"v0.2",
			"v0.2.b-1",
		},
	}

	for i, test := range table {
		result := cmd.GetBuilderVersion(test.Tags, test.Version)
		if result != test.Expected {
			t.Errorf("Test %d failed. \"%s\" != \"%s\" ", i, result, test.Expected)
		}
	}
}
