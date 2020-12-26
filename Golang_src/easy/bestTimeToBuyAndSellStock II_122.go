package easy

// Use greedy algorithm:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 96.90%
// Memory Usage: 3.1 MB, less than 15.75%
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	result := 0

	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			result += prices[i] - prices[i-1]
		}
	}

	return result
}
