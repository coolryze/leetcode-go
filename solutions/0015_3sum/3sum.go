package threesum

import "sort"

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}

		n1 := nums[i]
		if i > 0 && n1 == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 < 0 {
				l++
			} else if n1+n2+n3 > 0 {
				r--
			} else {
				res = append(res, []int{n1, n2, n3})
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
			}
		}
	}

	return res
}
