package twosum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for curIdx, num := range nums {
		if preIdx, ok := m[target-num]; ok {
			return []int{preIdx, curIdx}
		} else {
			m[num] = curIdx
		}
	}

	return []int{}
}
