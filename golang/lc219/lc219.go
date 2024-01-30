package lc219

func containsNearbyDuplicate(nums []int, k int) bool {
	alreadySeen := make(map[int]int)
	for i, val := range nums {
		if index, exists := alreadySeen[val]; exists {
			if index+k >= i {
				return true
			}
		}
		alreadySeen[val] = i
	}
	return false
}
