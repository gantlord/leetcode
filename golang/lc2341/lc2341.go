package lc2341

import "sort"

func numberOfPairs(nums []int) []int {
	sort.Ints(nums)
	counts := map[int]int{}
	for _, n := range nums {
		counts[n]++
	}
	pairs, spares := 0, 0
	for _, v := range counts {
		pairs += v / 2
		spares += v % 2
	}
	return []int{pairs, spares}
}
