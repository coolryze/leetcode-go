package sumofleftleaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	}

	lv := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		lv = root.Left.Val
	} else {
		lv = sumOfLeftLeaves(root.Left)
	}

	rv := sumOfLeftLeaves(root.Right)

	return lv + rv
}
