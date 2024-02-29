package lc2670

func distinctDifferenceArray(nums []int) []int {
	l := len(nums)
	fromLeft, fromRight := map[int]bool{nums[0]: true}, map[int]bool{nums[l-1]: true}
	distinctLeft, distinctRight := make([]int, l), make([]int, l)
	for i := 0; i < l; i++ {
		fromLeft[nums[i]] = true
		distinctLeft[i] = len(fromLeft)
		if i > 0 {
			distinctRight[l-i-1] = len(fromRight)
			fromRight[nums[i]] = true
		}
	}
	result := make([]int, l)
	for i := 0; i < l; i++ {
		result[i] = distinctLeft[i] - distinctRight[i]
	}
	return result
}
