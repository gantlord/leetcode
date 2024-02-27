package lc1614

func maxDepth(s string) int {
	maxOpen := 0
	open := 0
	for _, r := range s {
		if r == '(' {
			open++
		} else if r == ')' {
			open--
		}
		maxOpen = max(open, maxOpen)
	}
	return maxOpen
}
