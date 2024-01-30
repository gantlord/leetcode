package lc383

func canConstruct(ransomNote string, magazine string) bool {
	var runeStore = make(map[rune]int)
	for _, r := range magazine {
		if _, exists := runeStore[r]; !exists {
			runeStore[r] = 1
		} else {
			runeStore[r]++
		}
	}
	for _, r := range ransomNote {
		if _, exists := runeStore[r]; !exists {
			return false
		} else {
			runeStore[r]--
			if runeStore[r] < 0 {
				return false
			}
		}
	}
	return true
}
