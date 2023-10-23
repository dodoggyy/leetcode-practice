package easy

import "math"

// Use 2 stacks:
// Time Complexity: O(1)
// Space Complexity:O(n)
// Runtime: 17 ms, faster than 52.23%
// Memory Usage: 7.70 MB, less than 40.41%
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
		if val <= this.GetMin() {
			this.min = append(this.min, val)
		}
	}
}

func (this *MinStack) Pop() {
	if this.GetMin() == this.Top() {
		this.min = this.min[:len(this.min)-1]
	}
	this.data = this.data[:len(this.data)-1]

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

// Use 1 stack:
// Time Complexity: O(1)
// Space Complexity:O(n)
// Runtime: 10 ms, faster than 94.01%
// Memory Usage: 8.22 MB, less than 6.16%
type MinStack2 struct {
	data []int
	min  int
}

func ConstructorMinStack2() MinStack2 {
	return MinStack2{
		data: []int{},
		min:  math.MaxInt32,
	}
}

func (this *MinStack2) Push2(val int) {
	if val <= this.min {
		this.data = append(this.data, this.min)
		this.min = val
	}
	this.data = append(this.data, val)
}

func (this *MinStack2) Pop2() {
	if this.Top2() == this.GetMin2() {
		this.data = this.data[:len(this.data)-1]
		this.min = this.Top2()
		this.data = this.data[:len(this.data)-1]
	} else {
		this.data = this.data[:len(this.data)-1]
	}
}

func (this *MinStack2) Top2() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack2) GetMin2() int {
	return this.min
}
