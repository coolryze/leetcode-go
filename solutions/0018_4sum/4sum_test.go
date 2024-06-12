package foursum

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	nums   []int
	target int
}

type output struct {
	expected [][]int
}

func TestFourSum(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{1, 0, -1, 0, -2, 2}, 0},
			output{[][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		},
		testcase{
			input{[]int{2, 2, 2, 2, 2}, 8},
			output{[][]int{{2, 2, 2, 2}}},
		},
	}

	for _, tc := range tcs {
		got := fourSum(tc.input.nums, tc.input.target)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
