package pkg_test

import (
	"testing"

	"github.com/alauda/gitversion/pkg"
)

func TestPatchGit(t *testing.T) {

	type TestCase struct {
		Tags     []string
		Args     []string
		Expected string
	}

	table := []TestCase{
		{ // without prefix, adding v0.1-b.2 tag to make sure it will not affect
			[]string{"some", "stable-aaa", "v0.1-b.2", "v0.1.1"},
			[]string{"v0.1"},
			"v0.1.2",
		},
		{ // with minor suffix
			[]string{"some", "stable-aaa", "v0.1-b.2", "v0.1.1", "v0.1.5"},
			[]string{"v0.1-b"},
			"v0.1-b.3",
		},
		{ // with minor suffix and a dot
			[]string{"some", "stable-aaa", "v0.1-b.1", "v0.1-b.2", "v0.1-b.5", "v0.1.1", "v0.1.5"},
			[]string{"v0.1-b."},
			"v0.1-b.6",
		},
		{ // with empty after filter
			[]string{"some", "stable-aaa"},
			[]string{"v0.1-b."},
			"v0.1-b.0",
		},
	}

	for i, test := range table {
		result, err := pkg.PatchVersion(test.Args, test.Tags)
		if err != nil {
			t.Errorf("Test %d failed. should not return error: %v", i, err)
			continue
		}
		if result != test.Expected {
			t.Errorf("Test %d failed. \"%s\" != \"%s\" ", i, result, test.Expected)
		}
	}
}
