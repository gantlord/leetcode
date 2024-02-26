package lc1832

func checkIfPangram(sentence string) bool {
	seen := map[rune]bool{}
	for _, b := range sentence {
		seen[b] = true
		if len(seen) == 26 {
			return true
		}
	}
	return false
}
