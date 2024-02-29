package lc1716

func totalMoney(n int) int {
	weeks := n / 7
	days := n % 7
	return (((days + 1) * days) / 2) + ((max(weeks-1, 0)*weeks)/2)*7 + weeks*28 + days*weeks
}
