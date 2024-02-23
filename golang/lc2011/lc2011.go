package lc2011

func finalValueAfterOperations(operations []string) int {
	x := 0
	for _, o := range operations {
		switch o {
		case "++X":
			fallthrough
		case "X++":
			x++
		case "--X":
			fallthrough
		case "X--":
			x--
		}
	}
	return x
}
