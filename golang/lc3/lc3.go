package lc3

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	for i := 0; i < len(s); i++ {
		seen := map[byte]int{}
		l := 0
		for j := i; j < len(s); j++ {
			if index, exists := seen[s[j]]; exists {
				i = index
				break
			} else {
				seen[s[j]] = j
				l++
				maxLen = max(l, maxLen)
			}

		}
	}
	return maxLen
}
