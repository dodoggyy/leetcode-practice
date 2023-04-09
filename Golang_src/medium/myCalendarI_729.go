package medium

import "sort"

// Use brute force:
// Time Complexity: O(n^2)
// Space Complexity:O(n)
// Runtime: 106 ms, faster than 33.87%
// Memory Usage: 8 MB, less than 32.26%
type MyCalendar struct {
	Lists [][]int
}

func ConstructorMyCalendar() MyCalendar {
	return MyCalendar{
		Lists: [][]int{},
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	valid := this.isValid(start, end)
	if valid {
		this.Lists = append(this.Lists, []int{start, end})
	}

	return valid
}

func (this *MyCalendar) sort() {
	sort.Slice(this.Lists, func(i, j int) bool {
		return this.Lists[i][0] < this.Lists[j][0]
	})
}

func (this *MyCalendar) isValid(start, end int) bool {
	lists := this.Lists
	for _, list := range lists {
		if list[0] < end && start < list[1] {
			return false
		}
	}

	return true
}
