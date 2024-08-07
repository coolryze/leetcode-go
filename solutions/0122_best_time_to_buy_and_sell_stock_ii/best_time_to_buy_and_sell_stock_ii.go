package besttimetobuyandsellstockii

func maxProfit(prices []int) int {
	res := 0

	for i := 1; i < len(prices); i++ {
		res += max(0, prices[i]-prices[i-1])
	}

	return res
}
