package lc205

import "strings"

func isomorphify(s string) string {
	runeStore := make(map[rune]rune)
	runeIndex := 0
	f := func(r rune) rune {
		if _, exists := runeStore[r]; !exists {
			runeStore[r] = '0' + rune(runeIndex)
			runeIndex++
		}
		return runeStore[r]
	}
	return strings.Map(f, s)
}

func isIsomorphic(s string, t string) bool {
	sDescriptor := isomorphify(s)
	tDescriptor := isomorphify(t)
	return sDescriptor == tDescriptor
}
