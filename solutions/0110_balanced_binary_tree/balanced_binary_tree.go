package balancedbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var getHeight func(*TreeNode) int
	getHeight = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		lh := getHeight(node.Left)
		if lh == -1 {
			return -1
		}
		rh := getHeight(node.Right)
		if rh == -1 {
			return -1
		}

		if lh-rh > 1 || rh-lh > 1 {
			return -1
		}
		return max(lh, rh) + 1
	}

	return getHeight(root) != -1
}
