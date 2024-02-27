package lc2103

func countPoints(rings string) int {
	rods := make([]map[byte]bool, 10)
	for i := 0; i < 10; i++ {
		rods[i] = map[byte]bool{}
	}
	for i := 0; i < len(rings); i += 2 {
		rods[rings[i+1]-'0'][rings[i]] = true
	}
	count := 0
	for _, m := range rods {
		if len(m) == 3 {
			count++
		}
	}
	return count
}
