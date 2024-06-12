package threesum

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	nums []int
}

type output struct {
	expected [][]int
}

func TestThreeSum(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{-1, 0, 1, 2, -1, -4}},
			output{[][]int{{-1, -1, 2}, {-1, 0, 1}}},
		},
		testcase{
			input{[]int{0, 1, 1}},
			output{[][]int{}},
		},
		testcase{
			input{[]int{0, 0, 0}},
			output{[][]int{{0, 0, 0}}},
		},
	}

	for _, tc := range tcs {
		got := threeSum(tc.input.nums)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
