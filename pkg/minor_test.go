package pkg_test

import (
	"fmt"
	"testing"

	"github.com/alauda/gitversion/pkg"
)

func TestBumpMinor(t *testing.T) {
	type TestCase struct {
		Name    string
		Current string
		Result  string
		Err     error
	}

	table := []TestCase{
		{
			"only current version",
			"v0.1",
			"v0.2",
			nil,
		},
		{
			"empty",
			"",
			"",
			fmt.Errorf(""),
		},
		{
			"wrong format",
			"0.1.1",
			"",
			fmt.Errorf(""),
		},
	}

	for i, test := range table {
		res, err := pkg.BumpMinor(test.Current)
		if err != nil && test.Err == nil {
			t.Errorf("Test %d \"%s\" failed. Non expected error: %v", i, test.Name, err)
		} else if res != test.Result {
			t.Errorf("Test %d \"%s\" failed. Expected result error: %v != %v", i, test.Name, res, test.Result)
		}
	}
}
