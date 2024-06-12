package foursumii

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	nums1 []int
	nums2 []int
	nums3 []int
	nums4 []int
}

type output struct {
	expected int
}

func TestFourSumCount(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}},
			output{2},
		},
		testcase{
			input{[]int{0}, []int{0}, []int{0}, []int{0}},
			output{1},
		},
	}

	for _, tc := range tcs {
		got := fourSumCount(tc.input.nums1, tc.input.nums2, tc.input.nums3, tc.input.nums4)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
