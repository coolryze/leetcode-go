package reversestringii

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
	k int
}

type output struct {
	expected string
}

func TestReverseStr(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{"abcdefg", 2},
			output{"bacdfeg"},
		},
		testcase{
			input{"abcd", 2},
			output{"bacd"},
		},
	}

	for _, tc := range tcs {
		got := reverseStr(tc.input.s, tc.input.k)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
