package findbottomlefttreevalue

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	res := 0
	maxDepth := 0

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, curDepth int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if maxDepth < curDepth {
				maxDepth = curDepth
				res = node.Val
			}
			return
		}

		if node.Left != nil {
			dfs(node.Left, curDepth+1)
		}
		if node.Right != nil {
			dfs(node.Right, curDepth+1)
		}
	}

	dfs(root, 1)

	return res
}

func findBottomLeftValue2(root *TreeNode) int {
	res := 0
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		ql := queue.Len()
		for i := 0; i < ql; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if i == 0 {
				res = node.Val
			}

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}

	return res
}
