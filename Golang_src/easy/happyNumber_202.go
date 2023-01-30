package easy

// Use slow fast pointer:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 3 ms, faster than 45.61%
// Memory Usage: 1.9 MB, less than 100.00%
func isHappy(n int) bool {
	helper := func(num int) int {
		sum := 0
		for num > 0 {
			tmp := num % 10
			sum += tmp * tmp
			num /= 10
		}

		return sum
	}

	slow, fast := n, helper(n)

	for fast != 1 && slow != fast {
		slow = helper(slow)
		fast = helper(helper(fast))
	}

	return fast == 1
}
