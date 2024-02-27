package lc2956

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	nums1Map := map[int]bool{}
	nums2Map := map[int]bool{}
	for _, n := range nums1 {
		nums1Map[n] = true
	}
	result := make([]int, 2)
	for _, n := range nums2 {
		nums2Map[n] = true
		if _, exists := nums1Map[n]; exists {
			result[1]++
		}
	}
	for _, n := range nums1 {
		if _, exists := nums2Map[n]; exists {
			result[0]++
		}
	}
	return result
}
