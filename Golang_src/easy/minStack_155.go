package easy

// Use min stack:
// Time Complexity: O(1)
// Space Complexity:O(n)
// Runtime: 16 ms, faster than 88.79%
// Memory Usage: 7.8 MB, less than 23.28%
type MinStack struct {
	data []int
	min  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: []int{},
		min:  []int{}}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	if len(this.min) == 0 {
		this.min = append(this.min, val)
	} else {
		if this.GetMin() > val {
			this.min = append(this.min, val)
		} else {
			this.min = append(this.min, this.GetMin())
		}
	}
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
