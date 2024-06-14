package reversewordsinastring

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

func TestReverseWords(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{"the sky is blue"},
			output{"blue is sky the"},
		},
		testcase{
			input{"  hello world  "},
			output{"world hello"},
		},
		testcase{
			input{"a good   example"},
			output{"example good a"},
		},
	}

	for _, tc := range tcs {
		got := reverseWords(tc.input.s)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
