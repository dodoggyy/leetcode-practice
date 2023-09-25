package medium

// Use monotonic stack:
// Time Complexity: O(1)
// Space Complexity:O(n)
// Runtime: 134 ms, faster than 94.25%
// Memory Usage: 8.14 MB, less than 47.70%
type StockSpanner struct {
	stack []int
	idx   []int
}

func ConstructorStockSpanner() StockSpanner {
	return StockSpanner{
		stack: []int{},
		idx:   []int{},
	}
}

func (this *StockSpanner) Next(price int) int {
	result := 1
	for len(this.stack) > 0 && price >= this.stack[len(this.stack)-1] {
		this.stack = this.stack[:len(this.stack)-1]
		result += this.idx[len(this.idx)-1]
		this.idx = this.idx[:len(this.idx)-1]
	}
	this.stack = append(this.stack, price)
	this.idx = append(this.idx, result)

	return result
}
