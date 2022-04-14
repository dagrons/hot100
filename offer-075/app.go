package offer075

func findKthLargest(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	x := nums[0]
	i, j := -1, len(nums)
	for i < j {
		i++
		for nums[i] > x {
			i++
		}
		j--
		for nums[j] < x {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if j+1 >= k {
		return findKthLargest(nums[:j+1], k)
	} else {
		return findKthLargest(nums[j+1:], k-j-1)
	}
}
