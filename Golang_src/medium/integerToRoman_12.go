package medium

// Use greedy:
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 74.56%
// Memory Usage: 3.17 MB, less than 72.78%
func intToRoman(num int) string {
	result := []byte{}

	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for i, v := range nums {
		for num >= v {
			num -= v
			result = append(result, symbols[i]...)
		}
		if num == 0 {
			break
		}
	}

	return string(result)
}

// Use greedy 2:
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 3.17 MB, less than 72.78%
func intToRoman2(num int) string {
	result := []byte{}

	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	idx := 0
	for idx < len(nums) {
		for num >= nums[idx] {
			num -= nums[idx]
			result = append(result, symbols[idx]...)
		}
		idx++
	}

	return string(result)
}
