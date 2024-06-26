package constructbinarytreefrompreorderandinordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	var rootIdx int
	for idx, v := range inorder {
		if v == rootVal {
			rootIdx = idx
		}
	}

	leftInorder := inorder[:rootIdx]
	rightInorder := inorder[rootIdx+1:]

	leftPreorder := preorder[1 : 1+len(leftInorder)]
	rightPreorder := preorder[1+len(leftInorder):]

	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rightPreorder, rightInorder)

	return root
}
