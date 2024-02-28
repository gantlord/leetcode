package lc2427

func commonFactors(a int, b int) int {
	count := 0
	minArg := min(a, b)
	for i := 1; i <= minArg; i++ {
		if a%i == 0 && b%i == 0 {
			count++
		}
	}
	return count
}
