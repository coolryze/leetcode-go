package twosum

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
	expected []int
}

func TestTwoSum(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{2, 7, 11, 15}, 9},
			output{[]int{0, 1}},
		},
		testcase{
			input{[]int{3, 2, 4}, 6},
			output{[]int{1, 2}},
		},
		testcase{
			input{[]int{3, 3}, 6},
			output{[]int{0, 1}},
		},
	}

	for _, tc := range tcs {
		got := twoSum(tc.input.nums, tc.input.target)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
