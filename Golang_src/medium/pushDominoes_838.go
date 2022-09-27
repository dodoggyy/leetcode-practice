package medium

// Use simulate:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 42 ms, faster than 16.67%
// Memory Usage: 8.6 MB, less than 58.33%
func pushDominoes(dominoes string) string {
	// find range of '.'
	// then judge 3 cases of range of dots :
	// case left == right, then dots will change to same direction
	// case left = 'R' && right = 'L', then dots will change separate into middle
	// case left = 'L' && right = 'R', then dots will no change

	result := []rune(dominoes)
	cur, size, left := 0, len(result), rune('L')
	for cur < size {
		tmp := cur
		for tmp < size && result[tmp] == '.' {
			tmp++
		}
		right := rune('R')
		if tmp < size {
			right = result[tmp]
		}

		if left == right {
			for cur < tmp {
				result[cur] = right
				cur++
			}
		} else if left == 'R' && right == 'L' {
			k := tmp - 1
			for cur < k {
				result[cur] = 'R'
				result[k] = 'L'
				cur++
				k--
			}
		}
		left = right
		cur = tmp + 1
	}

	return string(result)
}
