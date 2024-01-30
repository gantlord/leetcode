package lc125

import (
	"strings"
	"unicode"
)

func alphaNumericMapper(r rune) rune {
	if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
		return -1
	}
	return r
}

func isPalindrome(s string) bool {
	sLower := strings.ToLower(s)
	sClean := strings.Map(alphaNumericMapper, sLower)
	i, j := 0, len(sClean)-1
	for i < j {
		if sClean[i] != sClean[j] {
			return false
		}
		i++
		j--
	}
	return true
}
