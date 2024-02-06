package lc134

func canCompleteCircuit(gas []int, cost []int) int {
	var start, totalGas, totalCost, gasLeft int
	for i, gasAvailable := range gas {
		gasLeft += gasAvailable - cost[i]
		totalGas += gasAvailable
		totalCost += cost[i]
		if gasLeft < 0 {
			start = i + 1
			gasLeft = 0
		} else if i == len(gas)-1 && totalCost <= totalGas {
			return start
		}
	}
	return -1
}
