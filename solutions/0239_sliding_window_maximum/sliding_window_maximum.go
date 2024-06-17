package slidingwindowmaximum

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0, len(nums)-k+1)
	dqueue := make([]int, 0, k)

	for i := 0; i < len(nums); i++ {
		if len(dqueue) > 0 && dqueue[0] < i-k+1 {
			dqueue = dqueue[1:]
		}

		for len(dqueue) > 0 && nums[dqueue[len(dqueue)-1]] < nums[i] {
			dqueue = dqueue[:len(dqueue)-1]
		}

		dqueue = append(dqueue, i)

		if i >= k-1 {
			res = append(res, nums[dqueue[0]])
		}
	}

	return res
}
