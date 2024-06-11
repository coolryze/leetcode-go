package validanagram

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
	t string
}

type output struct {
	expected bool
}

func TestIsAnagram(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{"anagram", "nagaram"},
			output{true},
		},
		testcase{
			input{"rat", "car"},
			output{false},
		},
	}

	for _, tc := range tcs {
		got := isAnagram(tc.input.s, tc.input.t)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
