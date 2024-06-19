package minimumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ld := minDepth(root.Left)
	rd := minDepth(root.Right)
	if root.Left == nil && root.Right != nil {
		return rd + 1
	}
	if root.Right == nil && root.Left != nil {
		return ld + 1
	}

	return min(ld, rd) + 1
}

func minDepth2(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for l := len(queue); l > 0; l = len(queue) {
		res++

		for i := l; i > 0; i-- {
			node := queue[0]
			if node.Left == nil && node.Right == nil {
				return res
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
	}

	return res
}
