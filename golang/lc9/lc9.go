package lc9

import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	xBinstring := strconv.FormatInt(int64(x), 10)
	for i := 0; i < len(xBinstring); i++ {
		if xBinstring[i] != xBinstring[len(xBinstring)-1-i] {
			return false
		}
	}
	return true
}
