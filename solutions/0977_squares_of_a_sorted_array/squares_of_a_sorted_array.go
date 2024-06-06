package squaresofasortedarray

func sortedSquares(nums []int) []int {
	l := len(nums)
	i, j, k := 0, l-1, l-1
	res := make([]int, l)

	for i <= j {
		is, js := nums[i]*nums[i], nums[j]*nums[j]
		if is < js {
			res[k] = js
			j--
		} else {
			res[k] = is
			i++
		}
		k--
	}

	return res
}
