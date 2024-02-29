package lc1207

import "sort"

func uniqueOccurrences(arr []int) bool {
	occurrences := map[int]bool{}
	sort.Ints(arr)
	cur := arr[0]
	count := 0
	for i, n := range arr {
		exists := false
		if n != cur {
			exists = occurrences[count]
			occurrences[count] = true
			count = 0
			cur = n
		}
		count++
		if i == len(arr)-1 {
			exists = exists || occurrences[count]
		}
		if exists {
			return false
		}
	}
	return true
}
