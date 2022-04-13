package i31

func NextPermutation(nums []int) {
	i := len(nums) - 1
	p, q := 0, len(nums)-1
	for ; i-1 >= 0 && nums[i-1] >= nums[i]; i-- {
	}
	if i >= 1 {
		j := len(nums) - 1
		for ; nums[j] <= nums[i-1]; j-- {
		}
		nums[i-1], nums[j] = nums[j], nums[i-1]
		p = i
	}
	for p < q {
		nums[p], nums[q] = nums[q], nums[p]
		p++
		q--
	}
	return
}
