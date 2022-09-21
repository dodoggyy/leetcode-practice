package medium

// Use Maintain Array Sum:
// Time Complexity: O(n+q)
// Space Complexity:O(n+q)
// Runtime: 184 ms, faster than 50.00%
// Memory Usage: 8.8 MB, less than 7.14%
func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	isEven := make([]bool, len(nums))
	keep := 0
	result := []int{}

	for i := range nums {
		if nums[i]%2 == 0 {
			isEven[i] = true
			keep += nums[i]
		}
	}

	for _, v := range queries {
		idx, val := v[1], v[0]

		nums[idx] += val

		if val%2 != 0 && !isEven[idx] { // odd + odd
			keep += nums[idx]
			isEven[idx] = true
		} else if val%2 == 0 && isEven[idx] { // even + even
			keep += val
		} else if val%2 != 0 && isEven[idx] { // odd + even
			keep -= nums[idx] - val
			isEven[idx] = false
		} else { // even + odd

		}

		result = append(result, keep)
	}

	return result
}

// Use optimize Maintain Array Sum:
// Time Complexity: O(n+q)
// Space Complexity:O(n)
// Runtime: 162 ms, faster than 64.29%
// Memory Usage: 8.4 MB, less than 7.14%
func sumEvenAfterQueries2(nums []int, queries [][]int) []int {
	keep := 0
	result := []int{}

	for i := range nums {
		if nums[i]%2 == 0 {
			keep += nums[i]
		}
	}

	for _, v := range queries {
		idx, val := v[1], v[0]

		if nums[idx]%2 == 0 {
			keep -= nums[idx]
		}

		nums[idx] += val

		if nums[idx]%2 == 0 {
			keep += nums[idx]
		}

		result = append(result, keep)
	}

	return result
}
