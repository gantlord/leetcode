package lc1309

import "strconv"

func getLetter(substring string) string {
	val, _ := strconv.ParseInt(substring, 10, 64)
	return string('a' + rune(val) - 1)
}

func freqAlphabets(s string) string {
	result := ""
	for i := 0; i < len(s); {
		if i < len(s)-2 && s[i+2] == '#' {
			result += getLetter(s[i : i+2])
			i += 3
		} else {
			result += getLetter(s[i : i+1])
			i++
		}
	}
	return result
}
