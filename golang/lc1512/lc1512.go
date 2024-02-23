package lc1512

func numIdenticalPairs(nums []int) int {
	count := 0
	for i, l := range nums {
		for j := i + 1; j < len(nums); j++ {
			if l == nums[j] {
				count++
			}
		}
	}
	return count
}
