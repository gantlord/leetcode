package lc2363

import "sort"

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	values := map[int]int{}
	for _, i := range append(items1, items2...) {
		weight := values[i[0]]
		values[i[0]] = weight + i[1]
	}
	result := make([][]int, len(values))
	i := 0
	for k, v := range values {
		result[i] = []int{k, v}
		i++
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	return result
}
