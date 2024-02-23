package lc452

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	candidate := intervals[0]
	result := [][]int{}
	for i := 1; i < len(intervals); i++ {
		if candidate[1] >= intervals[i][0] {
			candidate = []int{intervals[i][0], min(intervals[i][1], candidate[1])}
		} else {
			result = append(result, candidate)
			candidate = intervals[i]
		}
	}
	result = append(result, candidate)
	return result
}

func findMinArrowShots(points [][]int) int {
	points = merge(points)
	return len(points)
}
