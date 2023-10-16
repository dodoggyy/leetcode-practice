package medium

import "math"

// Use one pass iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 156 ms, faster than 49.21%
// Memory Usage: 9.83 MB, less than 68.48%
func canCompleteCircuit(gas []int, cost []int) int {
	minSpare := math.MaxInt32
	spare := 0
	result := 0
	for i := range gas {
		spare += gas[i] - cost[i]
		if spare <= minSpare {
			minSpare = spare
			result = i
		}
	}
	if spare < 0 {
		return -1
	}
	return (result + 1) % len(gas)
}
