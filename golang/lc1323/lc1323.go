package lc1323

func maximum69Number(num int) int {
	scale := 1
	biggest := 0
	working := num
	for working != 0 {
		digit := working % 10
		working /= 10
		if digit == 6 {
			biggest = scale
		}
		scale *= 10
	}
	num += biggest * 3
	return num
}
