package intersectionoftwoarrays

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
}

type output struct {
	expected []int
}

func TestIntersection(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{1, 2, 2, 1}, []int{2, 2}},
			output{[]int{2}},
		},
		testcase{
			input{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}},
			output{[]int{9, 4}},
		},
	}

	for _, tc := range tcs {
		got := intersection(tc.input.nums1, tc.input.nums2)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
