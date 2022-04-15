package i33

/*
	别人的思路:
	将一个数组一分为二，一定有一个数组是有序，一个无序的，
	对于有序的数组我们可以进行二分搜索，对于无序的部分可以再次一分而二
*/
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] >= nums[l] { // 左边有序
			if target <= nums[mid] && target >= nums[l] {
				r = mid
			} else {
				l = mid + 1
			}
		} else { // 左边无序
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid
			}
		}
	}
	if nums[l] == target {
		return l
	} else {
		return -1
	}
}
