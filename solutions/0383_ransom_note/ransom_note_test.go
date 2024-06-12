package ransomnote

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	ransomNote string
	magazine   string
}

type output struct {
	expected bool
}

func TestCanConstruct(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{"a", "b"},
			output{false},
		},
		testcase{
			input{"aa", "ab"},
			output{false},
		},
		testcase{
			input{"aa", "aab"},
			output{true},
		},
	}

	for _, tc := range tcs {
		got := canConstruct(tc.input.ransomNote, tc.input.magazine)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
