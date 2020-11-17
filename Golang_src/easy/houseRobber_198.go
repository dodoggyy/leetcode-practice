package easy

// Use DP:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 23.78%
func rob(nums []int) int {
	rob, notRob := 0, 0

	for _, num := range nums {
		cur := max(notRob+num, rob)
		notRob = rob
		rob = cur
	}

	return rob
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Use even & odd position:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 23.78%
func rob2(nums []int) int {
	robOdd, robEven := 0, 0

	for i := range nums {
		if i%2 == 0 {
			robEven = max(robEven+nums[i], robOdd)
		} else {
			robOdd = max(robOdd+nums[i], robEven)
		}
	}

	return max(robEven, robOdd)
}

// Use recursive with memorize:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.1 MB, less than 6.02%
func rob3(nums []int) int {
	size := len(nums)
	mem := make([]int, size)
	memView := make([]bool, size)
	return subRob(nums, mem, memView, size-1)
}

func subRob(nums, mem []int, memView []bool, index int) int {
	if index < 0 {
		return 0
	}
	if memView[index] == true {
		return mem[index]
	}

	mem[index] = max(subRob(nums, mem, memView, index-2)+nums[index], subRob(nums, mem, memView, index-1))

	memView[index] = true
	return mem[index]
}
