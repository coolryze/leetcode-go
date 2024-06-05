package binarysearch

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
	expected int
}

func TestBinarySearch(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{-1, 0, 3, 5, 9, 12}, 9},
			output{4},
		},
		testcase{
			input{[]int{-1, 0, 3, 5, 9, 12}, 2},
			output{-1},
		},
		testcase{
			input{[]int{-1, 0, 3, 5, 9, 12}, 13},
			output{-1},
		},
	}

	for _, tc := range tcs {
		got := search(tc.input.nums, tc.input.target)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%d, got=%d", tc.input, tc.output.expected, got)
		}
	}
}
