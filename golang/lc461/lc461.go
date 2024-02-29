package lc461

func hammingDistance(x int, y int) int {
	count := 0
	for x != 0 || y != 0 {
		xdigit, ydigit := x%2, y%2
		if xdigit != ydigit {
			count++
		}
		x /= 2
		y /= 2
	}
	return count
}
