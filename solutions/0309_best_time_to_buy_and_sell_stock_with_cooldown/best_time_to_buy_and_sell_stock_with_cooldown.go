package besttimetobuyandsellstockwithcooldown

func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i, _ := range dp {
		dp[i] = make([]int, 4)
	}
	dp[0][0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i], dp[i-1][3]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}

	return max(dp[len(prices)-1][1], dp[len(prices)-1][2], dp[len(prices)-1][3])
}
