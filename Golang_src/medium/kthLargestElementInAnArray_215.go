package medium

// Use quick select:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 32 ms, faster than 6.04%
// Memory Usage: 4 MB, less than 35.96%
func findKthLargest(nums []int, k int) int {
	return quickSelect(&nums, 0, len(nums)-1, k)
}

func quickSelect(nums *[]int, indexLow, indexHigh, k int) int {
	i := indexLow - 1

	for j := indexLow; j < indexHigh; j++ {
		if (*nums)[j] < (*nums)[indexHigh] {
			i++
			(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
		}
	}
	i++
	(*nums)[i], (*nums)[indexHigh] = (*nums)[indexHigh], (*nums)[i]

	count := indexHigh - i + 1
	if count == k {
		return (*nums)[i]
	} else if count > k {
		return quickSelect(nums, i+1, indexHigh, k)
	} else {
		return quickSelect(nums, indexLow, i - 1, k-count)
	}
}
