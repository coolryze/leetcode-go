package binarytreeinordertraversal

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left != nil {
			dfs(node.Left)
		}
		res = append(res, node.Val)
		if node.Right != nil {
			dfs(node.Right)
		}
	}

	dfs(root)
	return res
}

func inorderTraversal2(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	stack := list.New()
	cur := root

	for cur != nil || stack.Len() > 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			node := stack.Remove(stack.Back()).(*TreeNode)
			res = append(res, node.Val)
			cur = node.Right
		}
	}

	return res
}

func inorderTraversal3(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		e := stack.Remove(stack.Back())
		if e == nil {
			node := stack.Remove(stack.Back()).(*TreeNode)
			res = append(res, node.Val)
		} else {
			node := e.(*TreeNode)
			if node.Right != nil {
				stack.PushBack(node.Right)
			}
			stack.PushBack(node)
			stack.PushBack(nil)
			if node.Left != nil {
				stack.PushBack(node.Left)
			}
		}
	}

	return res
}
