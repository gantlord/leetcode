package lc1304

func sumZero(n int) []int {
	result := []int{}
	if n%2 != 0 {
		result = append(result, 0)
		n--
	}
	for i := 1; i <= n; i += 2 {
		result = append(result, []int{-i, i}...)
	}
	return result
}
