package lc12

func intToRoman(num int) string {
	numeralsInt := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	numeralsString := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	index := 0
	result := ""

	for index < len(numeralsInt) {
		if num >= numeralsInt[index] {
			num -= numeralsInt[index]
			result += numeralsString[index]
		} else {
			index++
		}
	}

	return result

}
