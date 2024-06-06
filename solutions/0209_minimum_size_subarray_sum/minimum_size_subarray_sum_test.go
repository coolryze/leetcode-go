package minimumsizesubarraysum

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	target int
	nums   []int
}

type output struct {
	expected int
}

func TestMinSubArrayLen(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{7, []int{2, 3, 1, 2, 4, 3}},
			output{2},
		},
		testcase{
			input{4, []int{1, 4, 4}},
			output{1},
		},
		testcase{
			input{11, []int{1, 1, 1, 1, 1, 1, 1, 1}},
			output{0},
		},
		testcase{
			input{11, []int{1, 2, 3, 4, 5}},
			output{3},
		},
	}

	for _, tc := range tcs {
		got := minSubArrayLen(tc.input.target, tc.input.nums)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%d, got=%d", tc.input, tc.output.expected, got)
		}
	}
}
