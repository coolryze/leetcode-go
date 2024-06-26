package pathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	var dfs func(*TreeNode, int) bool
	dfs = func(node *TreeNode, curSum int) bool {
		if node == nil {
			return false
		}

		curSum += node.Val
		if node.Left == nil && node.Right == nil {
			return curSum == targetSum
		}

		if dfs(node.Left, curSum) || dfs(node.Right, curSum) {
			return true
		}
		return false
	}

	return dfs(root, 0)
}

func hasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
