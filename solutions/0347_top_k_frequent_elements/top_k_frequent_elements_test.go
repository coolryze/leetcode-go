package topkfrequentelements

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

func TestTopKFrequent(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{1, 1, 1, 2, 2, 3}, 2},
			output{[]int{1, 2}},
		},
		testcase{
			input{[]int{1}, 1},
			output{[]int{1}},
		},
	}

	for _, tc := range tcs {
		got := topKFrequent(tc.input.nums, tc.input.k)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
