package validatebinarysearchtree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var check func(node *TreeNode, min, max int) bool
	check = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		if node.Val <= min || node.Val >= max {
			return false
		}

		return check(node.Left, min, node.Val) && check(node.Right, node.Val, max)
	}

	left := root.Left == nil || root.Left.Val < root.Val
	right := root.Right == nil || root.Right.Val > root.Val

	if left == false || right == false {
		return false
	}

	return check(root, math.MinInt, math.MaxInt)
}
