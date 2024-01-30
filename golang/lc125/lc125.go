package lc125

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	sLower := strings.ToLower(s)
	r, err := regexp.Compile("[^a-z0-9]+")
	if err != nil {
		return false
	}
	sClean := r.ReplaceAllString(sLower, "")
	sCleanReverseRunes := make([]rune, len(sClean))
	for i, c := range sClean {
		sCleanReverseRunes[len(sClean)-i-1] = c
	}
	return sClean == string(sCleanReverseRunes)
}
