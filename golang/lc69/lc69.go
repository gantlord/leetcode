package lc69

func mySqrt(x int) int {
	l, r := 0, 1<<16
	for l < r {
		m := l + (r-l)/2
		s := m * m
		if s == x {
			return m
		}
		if s > x {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	if l*l > x {
		l--
	}
	return l
}
