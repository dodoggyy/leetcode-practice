package hard

// Use 2 pass iterative:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 20 ms, faster than 31.74%
// Memory Usage: 6.09 MB, less than 99.19%
func candy(ratings []int) int {
	result := 0
	left := make([]int, len(ratings))

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i, v := range ratings {
		if i > 0 && v > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}

	right := 0
	for i := len(ratings) - 1; i >= 0; i-- {
		if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		result += max(left[i], right)
	}

	return result
}
