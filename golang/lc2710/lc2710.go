package lc2710

import "strings"

func removeTrailingZeros(num string) string {
	result := strings.TrimRight(num, "0")
	return result
}
