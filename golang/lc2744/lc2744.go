package lc2744

func reverse(s string) string {
	result := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}

func maximumNumberOfStringPairs(words []string) int {
	used := map[int]bool{}
	count := 0
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			usedI, usedJ := used[i], used[j]
			if !usedI && !usedJ && words[i] == reverse(words[j]) {
				count++
				used[j] = true
				used[i] = true
			}
		}
	}
	return count
}
