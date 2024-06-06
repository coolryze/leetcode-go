package squaresofasortedarray

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
	expected []int
}

func TestSquaresOfasortedarray(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{-4, -1, 0, 3, 10}},
			output{[]int{0, 1, 9, 16, 100}},
		},
		testcase{
			input{[]int{-7, -3, 2, 3, 11}},
			output{[]int{4, 9, 9, 49, 121}},
		},
	}

	for _, tc := range tcs {
		got := sortedSquares(tc.input.nums)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%d, got=%d", tc.input, tc.output.expected, got)
		}
	}
}
