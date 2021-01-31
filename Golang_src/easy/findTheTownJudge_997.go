package easy

// Use hashmap:
// Time Complexity: O(n^2)
// Space Complexity:O(n)
// Runtime: 8 ms, faster than 57.81%
// Memory Usage: 6 MB, less than 26.56%
func findJudge(N int, trust [][]int) int {
	if N == 1 && len(trust) == 0 {
		return 1
	}

	hashmap := map[int]int{}

	for _, v := range trust {
		if count, exist := hashmap[v[1]]; exist {
			hashmap[v[1]] = count + 1
		} else {
			hashmap[v[1]] = 1
		}
		if hashmap[v[1]] == (N - 1) {
			for i := range trust {
				if trust[i][0] == v[1] {
					return -1
				}
			}
			return v[1]
		}
	}

	return -1
}

// Use array count:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 100 ms, faster than 83.54%
// Memory Usage: 7.1 MB, less than 89.87%
func findJudge2(N int, trust [][]int) int {
	degree := make([]int, N+1)

	for _, v := range trust {
		degree[v[0]]--
		degree[v[1]]++
	}

	for i := 1; i < len(degree); i++ {
		if degree[i] == N-1 {
			return i
		}
	}

	return -1
}
