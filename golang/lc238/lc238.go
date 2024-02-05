package lc238

func productExceptSelf(nums []int) []int {
	leftProducts := make([]int, len(nums))
	rightProducts := make([]int, len(nums))
	leftProducts[0] = 1
	rightProducts[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		leftProducts[i] = leftProducts[i-1] * nums[i-1]
		rightProducts[len(nums)-i-1] = rightProducts[len(nums)-i] * nums[len(nums)-i]

	}
	for i := 0; i < len(nums); i++ {
		nums[i] = leftProducts[i] * rightProducts[i]
	}
	return nums
}
