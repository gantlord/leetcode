package lc657

func judgeCircle(moves string) bool {
	state := []int{0, 0}
	for _, m := range moves {
		switch m {
		case 'U':
			state[1]++
		case 'D':
			state[1]--
		case 'L':
			state[0]--
		case 'R':
			state[0]++
		}
	}
	return state[0] == 0 && state[1] == 0
}
