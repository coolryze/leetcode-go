package findbottomlefttreevalue

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
	expected int
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

func TestFindBottomLeftValue(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]interface{}{2, 1, 3}},
			output{1},
		},
		testcase{
			input{[]interface{}{1, 2, 3, 4, nil, 5, 6, nil, nil, 7}},
			output{7},
		},
		testcase{
			input{[]interface{}{1}},
			output{1},
		},
	}

	for _, tc := range tcs {
		got := findBottomLeftValue(arrayToTree(tc.input.nums))
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
