package lc2220

func minBitFlips(start int, goal int) int {
	count := 0
	for flips := goal ^ start; flips != 0; flips /= 2 {
		if flips%2 == 1 {
			count++
		}
	}
	return count
}
