package lc1812

func squareIsWhite(coordinates string) bool {
	offset := coordinates[0] - 'a' + coordinates[1] - '1'
	return offset%2 == 1
}
