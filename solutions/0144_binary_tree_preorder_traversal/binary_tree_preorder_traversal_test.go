package binarytreepreordertraversal

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	nums []interface{}
}

type output struct {
	expected []int
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

func TestPreorderTraversal(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]interface{}{1, nil, 2, 3}},
			output{[]int{1, 2, 3}},
		},
	}

	for _, tc := range tcs {
		got := preorderTraversal(arrayToTree(tc.input.nums))
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
