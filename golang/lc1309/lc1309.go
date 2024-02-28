package lc1309

func getLetter(substring []byte) byte {
	val := (substring[0]-'0')*10 + substring[1] - '0'
	return 'a' + val - 1
}

func freqAlphabets(s string) string {
	result := []byte{}
	buffer := make([]byte, 2)
	for i := 0; i < len(s); {
		if i < len(s)-2 && s[i+2] == '#' {
			buffer[0], buffer[1] = s[i], s[i+1]
			result = append(result, getLetter(buffer))
			i += 3
		} else {
			buffer[0], buffer[1] = '0', s[i]
			result = append(result, getLetter(buffer))
			i++
		}
	}
	return string(result)
}
