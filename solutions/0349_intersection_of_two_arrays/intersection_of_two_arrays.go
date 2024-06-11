package intersectionoftwoarrays

func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	m := make(map[int]struct{}, 0)

	for _, i := range nums1 {
		if _, ok := m[i]; !ok {
			m[i] = struct{}{}
		}
	}

	for _, i := range nums2 {
		if _, ok := m[i]; ok {
			delete(m, i)
			res = append(res, i)
		}
	}

	return res
}
