package lc1688

func numberOfMatches(n int) int {
	count := 0
	for n > 1 {
		count += n / 2
		n = n/2 + n%2
	}
	return count
}
