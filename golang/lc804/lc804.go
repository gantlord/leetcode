package lc804

func uniqueMorseRepresentations(words []string) int {
	morseArray := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	transforms := map[string]bool{}
	count := 0
	for _, w := range words {
		transform := ""
		for _, r := range w {
			transform += morseArray[r-'a']
		}
		if _, exists := transforms[transform]; !exists {
			count++
			transforms[transform] = true
		}
	}
	return count
}
