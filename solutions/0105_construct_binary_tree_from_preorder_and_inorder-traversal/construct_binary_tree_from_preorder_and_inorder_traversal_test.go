package constructbinarytreefrompreorderandinordertraversal

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	inorder   []int
	postorder []int
}

type output struct {
	expected []interface{}
}

func treeToLevelOrderArray(root *TreeNode) []interface{} {
	var res []interface{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			res = append(res, node.Val)
			queue = append(queue, node.Left, node.Right)
		} else {
			res = append(res, nil)
		}
	}

	for len(res) > 0 && res[len(res)-1] == nil {
		res = res[:len(res)-1]
	}

	return res
}

func TestBuildTree(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}},
			output{[]interface{}{3, 9, 20, nil, nil, 15, 7}},
		},
		testcase{
			input{[]int{-1}, []int{-1}},
			output{[]interface{}{-1}},
		},
	}

	for _, tc := range tcs {
		got := treeToLevelOrderArray(buildTree(tc.input.inorder, tc.input.postorder))
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
