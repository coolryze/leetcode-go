package validparentheses

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	s string
}

type output struct {
	expected bool
}

func TestIsValid(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{"()"},
			output{true},
		},
		testcase{
			input{"()[]{}"},
			output{true},
		},
		testcase{
			input{"(]"},
			output{false},
		},
	}

	for _, tc := range tcs {
		got := isValid(tc.input.s)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
