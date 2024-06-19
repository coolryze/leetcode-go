package binarytreepreordertraversal

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		res = append(res, node.Val)
		if node.Left != nil {
			dfs(node.Left)
		}
		if node.Right != nil {
			dfs(node.Right)
		}
	}

	dfs(root)
	return res
}

func preorderTraversal2(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		res = append(res, node.Val)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}

	return res
}

func preorderTraversal3(root *TreeNode) []int {
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
			if node.Left != nil {
				stack.PushBack(node.Left)
			}
			stack.PushBack(node)
			stack.PushBack(nil)
		}
	}

	return res
}
