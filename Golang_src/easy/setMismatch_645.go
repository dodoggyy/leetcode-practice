package easy

// Use Constant Space:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 48 ms, faster than 57.41%
// Memory Usage: 6.8 MB, less than 44.44%
func findErrorNums(nums []int) []int {
	for i := range nums {
		for nums[i] != (i+1) && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := range nums {
		if nums[i] != (i + 1) {
			return []int{nums[i], i + 1}
		}
	}
	return []int{}
}

// Use XOR:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 32 ms, faster than 67.74%
// Memory Usage: 6.6 MB, less than 83.87%
func findErrorNums2(nums []int) []int {
	xor := 0
	for _, v := range nums {
		xor ^= v
	}
	for i := 1; i <= len(nums); i++ {
		xor ^= i
	}

	lowBit := xor & -xor
	num1, num2 := 0, 0
	for _, v := range nums {
		if v&lowBit != 0 {
			num1 ^= v
		} else {
			num2 ^= v
		}
	}
	for i := 1; i <= len(nums); i++ {
		if i&lowBit != 0 {
			num1 ^= i
		} else {
			num2 ^= i
		}
	}

	for _, v := range nums {
		if v == num1 {
			return []int{num1, num2}
		}
	}

	return []int{num2, num1}
}
