package minimumsizesubarraysum

func minSubArrayLen(target int, nums []int) int {
	l := len(nums)
	res := l + 1
	i := 0
	sum := 0

	for j := 0; j < l; j++ {
		sum += nums[j]
		for sum >= target {
			sl := j - i + 1
			if sl < res {
				res = sl
			}
			sum -= nums[i]
			i++
		}
	}

	if res == l+1 {
		return 0
	} else {
		return res
	}
}
