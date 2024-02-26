package lc1342

func numberOfSteps(num int) int {
	result := 0
	for num != 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		result++
	}
	return result
}
