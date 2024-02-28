package lc3019

import "strings"

func countKeyChanges(s string) int {
	s = strings.ToLower(s)
	count := 0
	for i, r := range s[1:] {
		if r != rune(s[i]) {
			count++
		}
	}
	return count
}
