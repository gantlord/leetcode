package lc1221

func balancedStringSplit(s string) int {
	lCount, rCount, result := 0, 0, 0
	for _, c := range s {
		if c == 'L' {
			lCount++
		} else {
			rCount++
		}
		if lCount == rCount {
			result++
		}
	}
	return result
}
