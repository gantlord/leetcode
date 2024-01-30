package lc28

func strStr(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)
	window := " " + haystack[0:min(haystackLen, needleLen-1)]
	for i := 0; i < haystackLen; i++ {
		if haystackLen-i < needleLen {
			break
		}
		window = window[1:] + haystack[i+needleLen-1:i+needleLen]
		if window == needle {
			return i
		}
	}
	return -1
}
