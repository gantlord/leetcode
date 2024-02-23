package lc1431

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0
	for _, c := range candies {
		maxCandies = max(maxCandies, c)
	}
	result := make([]bool, len(candies))
	for i, c := range candies {
		if c+extraCandies >= maxCandies {
			result[i] = true
		}
	}
	return result
}
