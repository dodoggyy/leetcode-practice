package medium

// Use dp:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 88 ms, faster than 51.02%
// Memory Usage: 7.74 MB, less than 50.34%
func maxProfit(prices []int, fee int) int {
	// hold for max profit
	// sell for not hold profit
	hold := -prices[0] // intial condition for 1st
	sell := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(prices); i++ {
		tmp := sell
		sell = max(sell, hold+prices[i]-fee)
		hold = max(hold, tmp-prices[i])
	}

	return sell
}
