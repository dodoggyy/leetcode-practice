package medium

// Use hashset:
// m: the number of PopSmallest() method cal;
// n: the number AddBack()
// PopSmallest() Time Complexity: O((m+n)*logn)
// AddBack() Time Complexity: O(m*logn)
// Space Complexity:O(n)
// Runtime: 88 ms, faster than 51.02%
// Memory Usage: 7.74 MB, less than 50.34%
type SmallestInfiniteSet struct {
	set map[int]bool
}

func ConstructorSmallestInfiniteSet() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		set: map[int]bool{},
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	for i := 1; ; i++ {
		if !this.set[i] {
			this.set[i] = true
			return i
		}
	}
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	delete(this.set, num)
}
