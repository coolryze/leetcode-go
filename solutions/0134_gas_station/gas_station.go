package gasstation

func canCompleteCircuit(gas []int, cost []int) int {
	start := 0
	curSum := 0
	totalSum := 0

	for i := 0; i < len(gas); i++ {
		totalSum += gas[i] - cost[i]
		curSum += gas[i] - cost[i]
		if curSum < 0 {
			start = i + 1
			curSum = 0
		}
	}

	if totalSum < 0 {
		return -1
	}
	return start
}
