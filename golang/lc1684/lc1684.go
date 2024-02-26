package lc1684

import "strings"

func countConsistentStrings(allowed string, words []string) int {
	count := 0
	for _, w := range words {
		hasAll := true
		for _, r := range allowed {
			hasAll = hasAll && strings.ContainsRune(w, r)
		}
		if hasAll {
			count++
		}
	}
	return count
}
