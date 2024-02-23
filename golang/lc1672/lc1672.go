package lc1672

func maximumWealth(accounts [][]int) int {
	maxWealth := 0
	for _, c := range accounts {
		wealth := 0
		for _, a := range c {
			wealth += a
		}
		maxWealth = max(maxWealth, wealth)
	}
	return maxWealth
}
