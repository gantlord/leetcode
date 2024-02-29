package lc2215

func findDifference(nums1 []int, nums2 []int) [][]int {
	seen1, seen2 := map[int]bool{}, map[int]bool{}
	result := make([][]int, 2)
	result[0], result[1] = []int{}, []int{}
	for _, n := range nums1 {
		seen1[n] = true
	}
	for _, n := range nums2 {
		seen2[n] = true
	}
	for k := range seen1 {
		if !seen2[k] {
			result[0] = append(result[0], k)
		}
	}
	for k := range seen2 {
		if !seen1[k] {
			result[1] = append(result[1], k)
		}
	}
	return result
}
