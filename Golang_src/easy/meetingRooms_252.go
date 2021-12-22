package easy

import "sort"

// Given an array of meeting time intervals where intervals[i] = [starti, endi],
// determine if a person could attend all meetings.

// Example 1:
// Input: intervals = [[0,30],[5,10],[15,20]]
// Output: false

// Example 2:
// Input: intervals = [[7,10],[2,4]]
// Output: true

// Use sort + iterative:
// Time Complexity: O(nlogn)
// Space Complexity:O(1)
// Runtime: 8 ms, faster than 90.32%
// Memory Usage: 4.3 MB, less than 45.16%
func canAttendMeetings(intervals [][]int) bool {
	if len(intervals) == 0 {
		return true
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	cur := intervals[0]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < cur[1] {
			return false
		}
		cur = intervals[i]
	}

	return true
}
