package removealladjacentduplicatesinstring

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
	expected string
}

func TestRemoveDuplicates(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{"abbaca"},
			output{"ca"},
		},
		testcase{
			input{"azxxzy"},
			output{"ay"},
		},
	}

	for _, tc := range tcs {
		got := removeDuplicates(tc.input.s)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
