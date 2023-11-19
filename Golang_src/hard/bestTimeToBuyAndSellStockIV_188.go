package hard

import "math"

// Use dp:
// Time Complexity: O(n*k)
// Space Complexity:O(k)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.24 MB, less than 79.39%
func maxProfit2(k int, prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 0: sell
	// 1: buy

	// sell[i][j]=max{sell[i−1][j],buy[i−1][j−1]+price[i]}
	// buy[i][j]=max{buy[i−1][j],sell[i−1][j]−price[i]}

	dp := make([][2]int, k+2)
	for i := 1; i < k+2; i++ {
		dp[i][1] = math.MinInt32
	}

	dp[0][0] = math.MinInt32
	for i := range prices {
		for j := k + 1; j > 0; j-- {
			dp[j][0] = max(dp[j][0], dp[j][1]+prices[i])
			dp[j][1] = max(dp[j][1], dp[j-1][0]-prices[i])
		}
	}

	return dp[k+1][0]
}
