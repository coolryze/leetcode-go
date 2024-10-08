package binarysearch

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low < high {
		mid := low + (high-low)/2

		if nums[mid] > target {
			high = mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func search2(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2

		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
