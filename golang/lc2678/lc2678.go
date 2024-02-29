package lc2678

import "strconv"

func countSeniors(details []string) int {
	count := 0
	for _, n := range details {
		age, _ := strconv.Atoi(n[11:13])
		if age > 60 {
			count++
		}
	}
	return count
}
