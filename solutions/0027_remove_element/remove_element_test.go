package removeelement

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
	val  int
}

type output struct {
	expected int
}

func TestRemoveElement(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{3, 2, 2, 3}, 3},
			output{2},
		},
		testcase{
			input{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2},
			output{5},
		},
	}

	for _, tc := range tcs {
		got := removeElement(tc.input.nums, tc.input.val)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%d, got=%d", tc.input, tc.output.expected, got)
		}
	}
}
