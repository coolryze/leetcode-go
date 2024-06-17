package slidingwindowmaximum

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
	k    int
}

type output struct {
	expected []int
}

func TestMaxSlidingWindow(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3},
			output{[]int{3, 3, 5, 5, 6, 7}},
		},
	}

	for _, tc := range tcs {
		got := maxSlidingWindow(tc.input.nums, tc.input.k)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
