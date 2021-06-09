package medium

import "math/rand"

// Use Fisher-Yates algorithm:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 84 ms, faster than 24.56%
// Memory Usage: 10.5 MB, less than 5.26%
type Solution struct {
	arr []int
}

func Constructor384(nums []int) Solution {
	return Solution{nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.arr
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	size := len(this.arr)
	result := make([]int, size)
	copy(result, this.arr)

	for i := size - 1; i >= 0; i-- {
		randIdx := rand.Intn(i + 1)
		result[i], result[randIdx] = result[randIdx], result[i]
	}

	return result
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
