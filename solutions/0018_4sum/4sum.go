package foursum

import "sort"

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		n1 := nums[i]
		if n1 > target && (n1 > 0 || target > 0) {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			n2 := nums[j]
			if n1+n2 > target && (n1+n2 > 0 || target > 0) {
				break
			}
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			l, r := j+1, len(nums)-1
			for l < r {
				n3, n4 := nums[l], nums[r]
				if n1+n2+n3+n4 < target {
					l++
				} else if n1+n2+n3+n4 > target {
					r--
				} else {
					res = append(res, []int{n1, n2, n3, n4})
					for l < r && nums[l] == n3 {
						l++
					}
					for l < r && nums[r] == n4 {
						r--
					}
				}
			}
		}
	}

	return res
}
