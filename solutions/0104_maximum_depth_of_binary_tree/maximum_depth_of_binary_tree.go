package maximumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func maxDepth2(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for l := len(queue); l > 0; l = len(queue) {
		for i := l; i > 0; i-- {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}

		res++
	}

	return res
}
