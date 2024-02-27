package lc1021

func removeOuterParentheses(s string) string {
	result := []rune{}
	open := 0
	for _, r := range s {
		if r == '(' && open == 0 {
			open++
			continue
		} else if r == ')' && open == 1 {
			open--
			continue
		}
		if r == '(' {
			open++
		}
		if r == ')' {
			open--
		}
		result = append(result, r)
	}
	return string(result)
}
