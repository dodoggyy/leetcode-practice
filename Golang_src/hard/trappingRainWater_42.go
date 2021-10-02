package hard

// Use two pointer:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 8 ms, faster than 38.59%
// Memory Usage: 3.9 MB, less than 48.00%
func trap(height []int) int {
	left, right := 0, len(height)-1
	result := 0
	leftMax, rightMax := 0, 0

	for left < right {
		// check left max & update
		if height[left] >= leftMax {
			leftMax = height[left]
		}
		// check right max & update
		if height[right] >= rightMax {
			rightMax = height[right]
		}

		if height[left] < height[right] {
			result += leftMax - height[left]
			left++
		} else {
			result += rightMax - height[right]
			right--
		}
	}

	return result
}
