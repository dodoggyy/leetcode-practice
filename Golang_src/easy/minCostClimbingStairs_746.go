package easy

// Use DP:
// f(n) = min(f(n-1),f(n-2)) + cost(n)
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 5 ms, faster than 66.04%
// Memory Usage: 2.9 MB, less than 87.08%
func minCostClimbingStairs(cost []int) int {
	pre, cur := cost[0], cost[1]

	for i := 2; i < len(cost); i++ {
		next := min(pre, cur) + cost[i]
		pre = cur
		cur = next
	}

	return min(pre, cur)
}
