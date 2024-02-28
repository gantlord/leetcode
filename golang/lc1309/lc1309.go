package lc1309

func getLetter(substring []byte) byte {
	val := (substring[0]-'0')*10 + substring[1] - '0'
	return 'a' + val - 1
}

func freqAlphabets(s string) string {
	result := make([]byte, 0, 1000)
	buffer := make([]byte, 2)
	for i := 0; i < len(s); {
		var letter byte
		if i < len(s)-2 && s[i+2] == '#' {
			buffer[0], buffer[1] = s[i], s[i+1]
			letter = getLetter(buffer)
			i += 3
		} else {
			buffer[0], buffer[1] = '0', s[i]
			letter = getLetter(buffer)
			i++
		}
		result = append(result, letter)
	}
	return string(result)
}
