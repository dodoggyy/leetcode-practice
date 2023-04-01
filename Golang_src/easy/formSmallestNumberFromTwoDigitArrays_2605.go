package easy

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 4 ms, faster than 100%
// Memory Usage: 2.1 MB, less than 100%
func minNumber(nums1 []int, nums2 []int) int {
	cnt1, cnt2 := make([]int, 10), make([]int, 10)

	for _, v := range nums1 {
		cnt1[v]++
	}

	for _, v := range nums2 {
		cnt2[v]++
	}

	min1, min2 := 0, 0

	for i := 1; i < 10; i++ {
		if cnt1[i] > 0 && cnt1[i] == cnt2[i] {
			return i
		}

		if min1 == 0 && cnt1[i] > 0 {
			min1 = i
		}

		if min2 == 0 && cnt2[i] > 0 {
			min2 = i
		}

	}

	if min1 < min2 {
		return min1*10 + min2
	}
	return min2*10 + min1
}
