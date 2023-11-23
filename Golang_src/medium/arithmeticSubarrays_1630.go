package medium

// Use iterative:
// m: query times
// n: length of nums
// Time Complexity: O(m*n)
// Space Complexity:O(n)
// Runtime: 19 ms, faster than 98.18%
// Memory Usage: 7.55 MB, less than 32.73%
func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	isArithmetic := func(arr []int) bool {
		size := len(arr)
		max, min := arr[0], arr[0]
		for _, v := range arr {
			if v > max {
				max = v
			}
			if v < min {
				min = v
			}
		}
		if max == min {
			return true
		}
		ratio := (max - min) / (size - 1)
		if (max - min) != ratio*(size-1) {
			return false
		}
		hashset := map[int]bool{}
		for _, v := range arr {
			hashset[v] = true
		}

		cur := size
		for cur > 0 {
			if !hashset[min] {
				return false
			}
			min += ratio
			cur--
		}
		return true
	}

	result := []bool{}

	for i := range l {
		result = append(result, isArithmetic(nums[l[i]:r[i]+1]))
	}

	return result
}
