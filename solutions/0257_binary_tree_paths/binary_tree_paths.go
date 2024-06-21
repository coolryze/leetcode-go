package binarytreepaths

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	res := []string{}

	var getPath func(*TreeNode, string)
	getPath = func(node *TreeNode, curPath string) {
		if node.Left == nil && node.Right == nil {
			curPath += strconv.Itoa(node.Val)
			res = append(res, curPath)
			return
		}

		curPath = curPath + strconv.Itoa(node.Val) + "->"
		if node.Left != nil {
			getPath(node.Left, curPath)
		}
		if node.Right != nil {
			getPath(node.Right, curPath)
		}
	}

	getPath(root, "")

	return res
}
