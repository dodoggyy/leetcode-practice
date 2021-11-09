package easy

// Use two pointers
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 93.43%
// Memory Usage: 3 MB, less than 98.36%
func twoSumII(numbers []int, target int) []int {
	size := len(numbers)
	left, right := 0, size-1

	for left < right {
		current := numbers[left] + numbers[right]
		if current < target {
			left++
		} else if current > target {
			right--
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return []int{}
}
