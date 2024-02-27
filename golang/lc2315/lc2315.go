package lc2315

func countAsterisks(s string) int {
	barCount := 0
	asteriskCount := 0
	for _, r := range s {
		if r == '|' {
			barCount++
		}
		if r == '*' && barCount%2 == 0 {
			asteriskCount++
		}
	}
	return asteriskCount
}
