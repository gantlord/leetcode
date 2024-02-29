package lc1941

func areOccurrencesEqual(s string) bool {
	byteCount := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if count, exists := byteCount[s[i]]; exists {
			byteCount[s[i]] = count + 1
		} else {
			byteCount[s[i]] = 1
		}
	}
	prev := byteCount[s[0]]
	for _, v := range byteCount {
		if prev != v {
			return false
		}
		prev = v
	}
	return true
}
