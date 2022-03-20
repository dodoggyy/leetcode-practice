package medium

// Use greedy:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 186 ms, faster than 40.38%
// Memory Usage: 7.6 MB, less than 50.00%
func minDominoRotations(tops []int, bottoms []int) int {
	// nums: record all element counts
	// as: only record tops element counts
	// bs: only record bottoms element counts
	as, bs, nums := make([]int, 7), make([]int, 7), make([]int, 7)

	for i := range tops {
		if tops[i] == bottoms[i] {
			// only record nums because tops & bottoms are the same, no need to rotate
			nums[tops[i]]++
		} else {
			nums[tops[i]]++
			nums[bottoms[i]]++
			as[tops[i]]++
			bs[bottoms[i]]++
		}
	}

	for k, v := range nums {
		// if nums counts equal the size, compare min rotate in as and bs array
		if v == len(tops) {
			return min(as[k], bs[k])
		}
	}

	return -1
}
