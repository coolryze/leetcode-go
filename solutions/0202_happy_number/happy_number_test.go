package happynumber

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	n int
}

type output struct {
	expected bool
}

func TestIsHappy(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{19},
			output{true},
		},
		testcase{
			input{2},
			output{false},
		},
	}

	for _, tc := range tcs {
		got := isHappy(tc.input.n)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
