package maximumbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	rootIdx, rootVal := 0, nums[0]
	for idx, val := range nums {
		if val > rootVal {
			rootIdx = idx
			rootVal = val
		}
	}

	leftNums := nums[:rootIdx]
	rightNums := nums[rootIdx+1:]

	root := &TreeNode{Val: rootVal}
	root.Left = constructMaximumBinaryTree(leftNums)
	root.Right = constructMaximumBinaryTree(rightNums)

	return root
}
