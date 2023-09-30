package medium

// Use backtracking:
// n: length of digits
// Time Complexity: O(4^n * n)
// Space Complexity:O(n)
// Runtime: 1 ms, faster than 83.74%
// Memory Usage: 2.02 MB, less than 46.80%
func letterCombinations(digits string) []string {
	hashmap := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	result := []string{}
	size := len(digits)

	if len(digits) == 0 {
		return result
	}

	keep := []byte{}
	var dfs func(idx int)
	dfs = func(idx int) {
		if len(keep) == size {
			tmp := make([]byte, len(keep))
			copy(tmp, keep)
			result = append(result, string(tmp))
		} else {
			letters := hashmap[digits[idx]]
			for i := range letters {
				keep = append(keep, letters[i])
				dfs(idx + 1)
				keep = keep[:len(keep)-1]
			}
		}
	}

	dfs(0)

	return result
}
