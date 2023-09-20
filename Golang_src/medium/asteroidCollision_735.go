package medium

// Use stack:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 11 ms, faster than 44.69%
// Memory Usage: 4.58 MB, less than 84.74%
func asteroidCollision(asteroids []int) []int {
	stack := []int{}

	for _, v := range asteroids {
		alive := true
		for alive && v < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {
			alive = stack[len(stack)-1] < -v
			if stack[len(stack)-1] <= -v {
				stack = stack[:len(stack)-1]
			}
		}
		if alive {
			stack = append(stack, v)
		}
	}

	return stack
}
