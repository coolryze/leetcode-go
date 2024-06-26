package pathsum

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
	expected bool
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

func TestHasPathSum(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1}, 22},
			output{true},
		},
		testcase{
			input{[]interface{}{1, 2, 3}, 5},
			output{false},
		},
		testcase{
			input{[]interface{}{}, 0},
			output{false},
		},
	}

	for _, tc := range tcs {
		got := hasPathSum(arrayToTree(tc.input.nums), tc.input.target)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
