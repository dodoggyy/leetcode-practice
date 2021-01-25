package easy

// Use greedy algorithm:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 97.28%
// Memory Usage: 3.1 MB, less than 17.07%
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0]
	result := 0

	for _, v := range prices {
		if v < minPrice {
			minPrice = v
		} else {
			result = max(result, v-minPrice)
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
