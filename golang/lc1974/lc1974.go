package lc1974

func abs(x rune) rune {
	if x < 0 {
		return -x
	}
	return x
}

func minTimeToType(word string) int {
	prev := 'a'
	time := rune(0)
	for _, r := range word {
		diff := abs(r - prev)
		time += min(diff, 26-diff)
		prev = r
	}
	return int(time) + len(word)
}
