package medium

// Use DP (0-1 knapsack problem):
// Time Complexity: O(n*target)
// Space Complexity:O(n*target)
// n : total of nums
// target: target number = total sum / 2
// Runtime: 16 ms, faster than 58.10%
// Memory Usage: 6.4 MB, less than 46.64%
func canPartition(nums []int) bool {
	// nums[i] == j, dp[i][j] = true
	// nums[i] < j, dp[i][j] = dp[i-1][j] || dp[i-1][j- nums[i]]

	sum, max := 0, 0
	// O(n)
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	if max > target {
		return false
	}

	size := len(nums)
	dp := make([][]bool, size)
	// O(n)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	// O(n)
	for i := 0; i < size; i++ {
		dp[i][0] = true
	}

	dp[0][nums[0]] = true
	// total: O(n*target)
	// O(n)
	for i := 1; i < size; i++ {
		cur := nums[i]
		// O(target)
		for j := 1; j <= target; j++ {
			if cur <= j {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-cur]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
		if dp[i][target] {
			return true
		}
	}

	return dp[size-1][target]
}

// Optimize DP:
// Time Complexity: O(n*target)
// Space Complexity:O(target)
// n : total of nums
// target: target number = total sum / 2
// Runtime: 12 ms, faster than 83.79%
// Memory Usage: 2.5 MB, less than 86.17%
func canPartition2(nums []int) bool {
	sum, max := 0, 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	if max > target {
		return false
	}

	dp := make([]bool, target+1)
	size := len(nums)

	dp[0] = true
	for i := 0; i < size; i++ {
		cur := nums[i]
		for j := target; j >= cur; j-- {
			dp[j] = dp[j] || dp[j-cur]
		}
	}

	return dp[target]
}
