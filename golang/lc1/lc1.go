package lc1

func twoSum(nums []int, target int) []int {
	foundSet := make(map[int]int)
	for i, val := range nums {
		complement := target - val
		if index, existsComplement := foundSet[complement]; existsComplement {
			return []int{index, i}
		}
		foundSet[val] = i
	}
	return nil
}
