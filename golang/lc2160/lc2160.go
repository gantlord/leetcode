package lc2160

import "sort"

func minimumSum(num int) int {
	d := make([]int, 4)
	for i := 0; i < 4; i++ {
		d[i] = num % 10
		num /= 10
	}
	sort.Ints(d)
	return (d[0]+d[1])*10 + d[2] + d[3]
}
