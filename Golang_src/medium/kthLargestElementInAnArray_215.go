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
		return quickSelect(nums, indexLow, i-1, k-count)
	}
}

// Use quick select 2:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 32 ms, faster than 5.33%
// Memory Usage: 3.9 MB, less than 42.13%
func findKthLargest2(nums []int, k int) int {
	var qSelect func(start, end, k int) int
	qSelect = func(start, end, k int) int {
		i := start - 1

		// O(n)
		for j := start; j < end; j++ {
			if nums[j] < nums[end] {
				i++
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		i++
		nums[end], nums[i] = nums[i], nums[end]

		cnt := end - i + 1
		if cnt == k {
			return nums[i]
		} else if cnt < k {
			return qSelect(start, i-1, k-cnt)
		} else {
			return qSelect(i+1, end, k)
		}
	}

	return qSelect(0, len(nums)-1, k)
}
