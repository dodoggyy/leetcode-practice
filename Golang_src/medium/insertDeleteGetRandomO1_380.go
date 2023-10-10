package medium

import "math/rand"

// Use hashmap:
// Time Complexity: O(1)
// Space Complexity:O(n)
// Runtime: 169 ms, faster than 39.03%
// Memory Usage: 54.84 MB, less than 37.10%
type RandomizedSet struct {
	hashmap map[int]int
	nums    []int
}

func ConstructorRandomizedSet() RandomizedSet {
	return RandomizedSet{
		hashmap: map[int]int{},
		nums:    []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hashmap[val]; ok {
		return false
	}
	this.hashmap[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.hashmap[val]; ok {
		valSwap := this.nums[len(this.nums)-1]
		this.nums[idx] = valSwap
		this.hashmap[valSwap] = idx
		delete(this.hashmap, val)
		this.nums = this.nums[:len(this.nums)-1]
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	randomIdx := rand.Intn(len(this.nums))
	return this.nums[randomIdx]
}
