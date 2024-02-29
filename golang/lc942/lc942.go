package lc942

func diStringMatch(s string) []int {
	l, r := 0, len(s)
	w := []rune(s)
	result := []int{}
	length := len(s) + 1
	for i := 0; i < length; i++ {
		c := w[i]
		if c == 'I' {
			result = append(result, l)
			l++
			if i == length-2 {
				w = append(w, 'D')
			}
		} else {
			result = append(result, r)
			r--
			if i == length-2 {
				w = append(w, 'I')
			}
		}
	}
	return result
}
