package lc1051

import "sort"

func heightChecker(heights []int) int {
	expected := append([]int{}, heights...)
	sort.Ints(expected)
	count := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != expected[i] {
			count++
		}
	}
	return count
}
