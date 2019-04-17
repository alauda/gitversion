package pkg_test

import (
	"testing"

	"github.com/alauda/gitversion/pkg"
)

func TestRCVersion(t *testing.T) {
	type TestCase struct {
		Name    string
		Current string
		Next    string
		Result  string
		Err     error
	}

	table := []TestCase{
		{
			"next and current are the same major minor",
			"v0.1.1",
			"v0.1",
			"v0.1-rc.0",
			nil,
		},
		{
			"next and current have different major minor",
			"v1.2.1",
			"v0.1",
			"v0.1-rc.0",
			nil,
		},
		{
			"current is RC",
			"v0.1-rc.1",
			"v0.1",
			"v0.1-rc.2",
			nil,
		},
		{
			"current is build",
			"v0.1-b.1",
			"v0.1",
			"v0.1-rc.0",
			nil,
		},
	}

	for i, test := range table {
		res, err := pkg.GenRC(test.Current, test.Next)
		if err != nil && test.Err == nil {
			t.Errorf("Test %d \"%s\" failed. Non expected error: %v", i, test.Name, err)
		} else if res != test.Result {
			t.Errorf("Test %d \"%s\" failed. Expected result error: %v != %v", i, test.Name, res, test.Result)
		}

	}
}
