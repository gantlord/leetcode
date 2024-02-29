package lc1475

func finalPrices(prices []int) []int {
	answer := make([]int, len(prices))
	for i := range prices {
		answer[i] = prices[i]
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				answer[i] = prices[i] - prices[j]
				break
			}
		}
	}
	return answer
}
