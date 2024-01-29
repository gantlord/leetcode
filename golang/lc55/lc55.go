package lc55

func canJump(nums []int) bool {
	maxReachableIndex := 0
	for i, elem := range nums {
		if maxReachableIndex < i {
			return false
		}
		maxReachableIndex = max(maxReachableIndex, i+elem)
	}
	return true
}
