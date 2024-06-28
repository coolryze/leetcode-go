package findmodeinbinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	res := []int{}

	maxCount := 1
	curCount := 1
	var prev *TreeNode
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)

		if prev != nil && node.Val == prev.Val {
			curCount++
		} else {
			curCount = 1
		}
		if curCount >= maxCount {
			if curCount > maxCount {
				res = []int{node.Val}
			} else {
				res = append(res, node.Val)
			}
			maxCount = curCount
		}

		prev = node
		dfs(node.Right)
	}

	dfs(root)

	return res
}
