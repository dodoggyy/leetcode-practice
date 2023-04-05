package medium

// Use one pass:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 2 ms, faster than 65.2%
// Memory Usage: 2.1 MB, less than 70.37%
func sortColors(nums []int) {
	idx0, idx1, idx2 := -1, -1, -1

	for _, v := range nums {
		switch v {
		case 0:
			idx0++
			idx1++
			idx2++
			nums[idx2] = 2
			nums[idx1] = 1
			nums[idx0] = 0
		case 1:
			idx1++
			idx2++
			nums[idx2] = 2
			nums[idx1] = 1
		case 2:
			idx2++
			nums[idx2] = 2
		}
	}
}
