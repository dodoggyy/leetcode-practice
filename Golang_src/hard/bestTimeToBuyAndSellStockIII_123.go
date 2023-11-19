package hard

// Use dp:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 95 ms, faster than 89.94%
// Memory Usage: 9.06 MB, less than 54.72%
func maxProfit(prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0

	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}
