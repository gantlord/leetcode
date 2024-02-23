package lc1281

func subtractProductAndSum(n int) int {
	s, p := 0, 1
	for n > 0 {
		digit := n % 10
		s += digit
		p *= digit
		n /= 10
	}
	return p - s
}
