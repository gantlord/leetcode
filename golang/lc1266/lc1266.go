package lc1266

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minTimeToVisitAllPoints(points [][]int) int {
	time := 0
	prev := points[0]
	for _, p := range points[1:] {
		time += max(abs(prev[0]-p[0]), abs(prev[1]-p[1]))
		prev = p
	}
	return time
}
