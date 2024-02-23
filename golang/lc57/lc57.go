package lc57

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	candidate := intervals[0]
	result := [][]int{}
	for i := 1; i < len(intervals); i++ {
		if candidate[1] >= intervals[i][0] {
			candidate = []int{candidate[0], max(intervals[i][1], candidate[1])}
		} else {
			result = append(result, candidate)
			candidate = intervals[i]
		}
	}
	result = append(result, candidate)
	return result
}

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	return merge(intervals)
}
