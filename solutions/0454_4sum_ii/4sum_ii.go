package foursumii

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count := 0
	m := map[int]int{}

	for _, a := range nums1 {
		for _, b := range nums2 {
			m[a+b]++
		}
	}

	for _, c := range nums3 {
		for _, d := range nums4 {
			if value, ok := m[-c-d]; ok {
				count += value
			}
		}
	}

	return count
}
