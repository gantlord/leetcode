package lc2960

func countTestedDevices(batteryPercentages []int) int {
	count := 0
	for _, d := range batteryPercentages {
		if d > count {
			count++
		}
	}
	return count
}
