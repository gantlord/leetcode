package lc3

func lengthOfLongestSubstring(s string) int {
	seen := map[rune]int{}
	maxLen, start := 0, 0
	for i, c := range s {
		if index, exists := seen[c]; exists && index >= start {
			start = index + 1
		}
		seen[c] = i
		if l := i - start + 1; l > maxLen {
			maxLen = l
		}
	}
	return maxLen
}
