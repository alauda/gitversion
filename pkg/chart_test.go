package pkg_test

import (
	"testing"

	"github.com/alauda/gitversion/pkg"
)

func TestCharVersion(t *testing.T) {
	type TestCase struct {
		Name    string
		Current string
		Next    string
		Result  string
		Err     error
	}

	table := []TestCase{
		{
			"only current version",
			"v0.1.0",
			"",
			"v0.1.1",
			nil,
		},
		{
			"next and current are the same",
			"v0.1.0",
			"v0.1.0",
			"v0.1.0",
			nil,
		},
		{
			"next is minor increased",
			"v0.1.0",
			"v0.2",
			"v0.2.0",
			nil,
		},
		{
			"next is minor is the same",
			"v0.1.1",
			"v0.1",
			"v0.1.2",
			nil,
		},
		{
			"next is greater than current",
			"v0.1.0",
			"v0.1.2",
			"v0.1.2",
			nil,
		},
		{
			"next is lesser than current",
			"v0.1.2",
			"v0.1.0",
			"v0.1.3",
			nil,
		},
	}

	for i, test := range table {
		res, err := pkg.GetNextChartVersion(test.Current, test.Next)
		if err != nil && test.Err == nil {
			t.Errorf("Test %d \"%s\" failed. Non expected error: %v", i, test.Name, err)
		} else if res != test.Result {
			t.Errorf("Test %d \"%s\" failed. Expected result error: %v != %v", i, test.Name, res, test.Result)
		}

	}
}
