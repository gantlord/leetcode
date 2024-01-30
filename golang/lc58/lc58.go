package lc58

func lengthOfLastWord(s string) int {
	lengthOfLastWord := 0
	seenSpace := true
	for _, elem := range s {
		if elem == ' ' {
			seenSpace = true
		} else {
			if seenSpace {
				lengthOfLastWord = 0
			}
			seenSpace = false
			lengthOfLastWord++
		}
	}
	return lengthOfLastWord
}
