package countcompletetreenodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ld, rd := 1, 1
	l, r := root.Left, root.Right
	for l != nil {
		l = l.Left
		ld++
	}
	for r != nil {
		r = r.Right
		rd++
	}

	if ld == rd {
		return 2<<(ld-1) - 1
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}
