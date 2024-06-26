package pathsumii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	var dfs func(*TreeNode, int, []int)
	dfs = func(node *TreeNode, sum int, nums []int) {
		if node.Left == nil && node.Right == nil {
			if node.Val == sum {
				temp := make([]int, len(nums))
				copy(temp, nums)
				res = append(res, temp)
			}
			return
		}

		if node.Left != nil {
			dfs(node.Left, sum-node.Val, append(nums, node.Left.Val))
		}
		if node.Right != nil {
			dfs(node.Right, sum-node.Val, append(nums, node.Right.Val))
		}
	}

	dfs(root, targetSum, []int{root.Val})
	return res
}
