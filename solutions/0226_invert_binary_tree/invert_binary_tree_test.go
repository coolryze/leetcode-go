package invertbinarytree

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

func treeToLevelOrderArray(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return res
}

func TestInvertTree(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]interface{}{4, 2, 7, 1, 3, 6, 9}},
			output{[]int{4, 7, 2, 9, 6, 3, 1}},
		},
		testcase{
			input{[]interface{}{2, 1, 3}},
			output{[]int{2, 3, 1}},
		},
	}

	for _, tc := range tcs {
		got := invertTree(arrayToTree(tc.input.nums))
		if !reflect.DeepEqual(treeToLevelOrderArray(got), tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
