package easy

// Use queue:
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 123 ms, faster than 48.73%
// Memory Usage: 7.9 MB, less than 89.86%
type RecentCounter struct {
	TimeSlots []int
}

func ConstructorRecentCounter() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	past := t - 3000
	this.TimeSlots = append(this.TimeSlots, t)

	for this.TimeSlots[0] < past {
		this.TimeSlots = this.TimeSlots[1:]
	}

	return len(this.TimeSlots)
}
