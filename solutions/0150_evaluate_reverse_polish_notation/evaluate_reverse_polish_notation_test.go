package evaluatereversepolishnotation

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	tokens []string
}

type output struct {
	expected int
}

func TestEvalRPN(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]string{"2", "1", "+", "3", "*"}},
			output{9},
		},
		testcase{
			input{[]string{"4", "13", "5", "/", "+"}},
			output{6},
		},
		testcase{
			input{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}},
			output{22},
		},
	}

	for _, tc := range tcs {
		got := evalRPN(tc.input.tokens)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
