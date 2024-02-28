package lc1704

import (
	"strings"
	"unicode"
)

func halvesAreAlike(s string) bool {
	vowels := "aeiou"
	count := 0
	midpoint := len(s) >> 1
	for i, r := range s {
		r = unicode.ToLower(r)
		if strings.ContainsRune(vowels, r) {
			if i < midpoint {
				count++
			} else {
				count--
			}
		}
	}
	return count == 0
}
