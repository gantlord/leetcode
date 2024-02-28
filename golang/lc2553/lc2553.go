package lc2553

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func separateDigits(nums []int) []int {
	result := make([]int, 0, 6000)
	for _, n := range nums {
		start := len(result)
		for n != 0 {
			result = append(result, n%10)
			n /= 10
		}
		reverse(result[start:])
	}
	return result
}
