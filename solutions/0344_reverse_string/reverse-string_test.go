package reversestring

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	s []byte
}

type output struct {
	expected []byte
}

func TestReverseString(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]byte{'h', 'e', 'l', 'l', 'o'}},
			output{[]byte{'o', 'l', 'l', 'e', 'h'}},
		},
		testcase{
			input{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}},
			output{[]byte{'h', 'a', 'n', 'n', 'a', 'H'}},
		},
	}

	for _, tc := range tcs {
		reverseString(tc.input.s)
		got := tc.input.s
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
