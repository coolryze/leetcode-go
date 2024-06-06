package spiralmatrixii

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	n int
}

type output struct {
	expected [][]int
}

func TestGenerateMatrix(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{3},
			output{[][]int{[]int{1, 2, 3}, []int{8, 9, 4}, []int{7, 6, 5}}},
		},
		testcase{
			input{1},
			output{[][]int{[]int{1}}},
		},
	}

	for _, tc := range tcs {
		got := generateMatrix(tc.input.n)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%d, got=%d", tc.input, tc.output.expected, got)
		}
	}
}
