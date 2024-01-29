package lc122

func maxProfit(prices []int) int {
	maxProfit := 0
	lastPrice := prices[0]
	for _, price := range prices[1:] {
		if price > lastPrice {
			maxProfit += price - lastPrice
		}
		lastPrice = price
	}
	return maxProfit
}
