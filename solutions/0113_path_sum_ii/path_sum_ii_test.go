package pathsumii

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	nums   []interface{}
	target int
}

type output struct {
	expected [][]int
}

func arrayToTree(nums []interface{}) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	root := &TreeNode{Val: nums[0].(int)}
	queue := []*TreeNode{root}
	index := 1

	for index < n {
		node := queue[0]
		queue = queue[1:]

		if index < n && nums[index] != nil {
			left := &TreeNode{Val: nums[index].(int)}
			node.Left = left
			queue = append(queue, left)
		}
		index++

		if index < n && nums[index] != nil {
			right := &TreeNode{Val: nums[index].(int)}
			node.Right = right
			queue = append(queue, right)
		}
		index++
	}

	return root
}

func TestPathSum(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1}, 22},
			output{[][]int{[]int{5, 4, 11, 2}, []int{5, 8, 4, 5}}},
		},
		testcase{
			input{[]interface{}{1, 2, 3}, 5},
			output{[][]int{}},
		},
		testcase{
			input{[]interface{}{}, 0},
			output{[][]int{}},
		},
	}

	for _, tc := range tcs {
		got := pathSum(arrayToTree(tc.input.nums), tc.input.target)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
