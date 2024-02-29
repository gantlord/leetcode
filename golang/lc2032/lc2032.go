package lc2032

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	seen, seen1, seen2, seen3 := map[int]int{}, map[int]int{}, map[int]int{}, map[int]int{}
	for _, n := range nums1 {
		seen1[n] = 1
		seen[n] = 1
	}
	for _, n := range nums2 {
		seen2[n] = 1
		seen[n] = 1
	}
	for _, n := range nums3 {
		seen3[n] = 1
		seen[n] = 1
	}
	result := []int{}
	for k := range seen {
		count := seen1[k] + seen2[k] + seen3[k]
		if count >= 2 {
			result = append(result, k)
		}
	}
	return result
}
