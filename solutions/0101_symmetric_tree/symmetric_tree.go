package symmetrictree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var dfs func(left *TreeNode, right *TreeNode) bool
	dfs = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
	}

	return dfs(root.Left, root.Right)
}

func isSymmetric2(root *TreeNode) bool {
	queue := []*TreeNode{root.Left, root.Right}

	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]

		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}

		queue = append(queue, left.Left, right.Right, left.Right, right.Left)
	}

	return true
}
