package lc1356

import "sort"

func numBits(n int) int {
	count := 0
	for n != 0 {
		if n%2 == 1 {
			count++
		}
		n /= 2
	}
	return count
}

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i int, j int) bool {
		iBits, jBits := numBits(arr[i]), numBits(arr[j])
		if iBits == jBits {
			return arr[i] < arr[j]
		} else {
			return iBits < jBits
		}
	})
	return arr
}
