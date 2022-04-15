package i34

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	return []int{down(nums, target), up(nums, target)}
}

func down(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) >> 1
		if target <= nums[mid] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}

func up(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r + 1) >> 1
		if target >= nums[mid] {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}
