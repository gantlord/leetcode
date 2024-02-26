package lc1773

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	result := 0
	for _, item := range items {
		match := false
		switch ruleKey {
		case "type":
			match = ruleValue == item[0]
		case "color":
			match = ruleValue == item[1]
		case "name":
			match = ruleValue == item[2]
		}
		if match {
			result++
		}
	}
	return result
}
