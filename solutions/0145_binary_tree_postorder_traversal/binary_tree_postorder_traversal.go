package binarytreepostordertraversal

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	res := []int{}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			dfs(node.Left)
		}
		if node.Right != nil {
			dfs(node.Right)
		}
		res = append(res, node.Val)
	}

	dfs(root)
	return res
}

func postorderTraversal2(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		res = append(res, node.Val)
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}

func postorderTraversal3(root *TreeNode) []int {
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
			stack.PushBack(node)
			stack.PushBack(nil)
			if node.Right != nil {
				stack.PushBack(node.Right)
			}
			if node.Left != nil {
				stack.PushBack(node.Left)
			}
		}
	}

	return res
}
