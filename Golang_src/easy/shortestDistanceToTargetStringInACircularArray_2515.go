package easy

import "math"

// Use two pass iterative:
// Time Complexity: O(2n) = O(n)
// Space Complexity:O(1)
// Runtime: 10 ms, faster than 51.2%
// Memory Usage: 6.9 MB, less than 24.49%
func closetTarget(words []string, target string, startIndex int) int {
	result := math.MaxInt32
	cnt := 0
	size := len(words)

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	cur := startIndex
	for words[cur] != target && cnt != size {
		cur = (cur + 1) % size
		cnt++
	}

	result = min(result, cnt)

	cur = startIndex
	cnt = 0
	for words[cur] != target && cnt != size {
		if cur-1 <= 0 {
			cur += size
		}
		cur = (cur - 1) % size
		cnt++
	}

	result = min(result, cnt)

	if result == size {
		return -1
	}
	return result
}

// Use one pass iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 7 ms, faster than 77.55%
// Memory Usage: 5.7 MB, less than 44.90%
func closetTarget2(words []string, target string, startIndex int) int {
	size := len(words)
	result := size

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	for i := range words {
		if words[i] == target {
			result = min(result, min(abs(i-startIndex), size-abs(i-startIndex)))
		}
	}

	if result == size {
		return -1
	}
	return result
}
