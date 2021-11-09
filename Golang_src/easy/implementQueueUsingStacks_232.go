package easy

// Use two stack:
// Time Complexity: O(1)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2 MB, less than 100.00%
type MyQueue struct {
	in, out []int
}

/** Initialize your data structure here. */
func ConstructorMyQueue() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) in2Out() {
	for len(this.in) > 0 {
		this.out = append(this.out, this.in[len(this.in)-1])
		this.in = this.in[:len(this.in)-1]
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}

	if len(this.out) == 0 {
		this.in2Out()
	}

	val := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]

	return val
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}

	if len(this.out) == 0 {
		this.in2Out()
	}

	return this.out[len(this.out)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.out) == 0 && len(this.in) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
