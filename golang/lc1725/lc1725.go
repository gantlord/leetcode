package lc1725

func countGoodRectangles(rectangles [][]int) int {
	maxLen := 0
	count := 0
	for _, r := range rectangles {
		len := min(r[0], r[1])
		if len > maxLen {
			maxLen = len
			count = 1
		} else if len == maxLen {
			count++
		}
	}
	return count
}
