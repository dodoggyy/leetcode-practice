package medium

// Use hashmap:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 168 ms, faster than 80.00%
// Memory Usage: 31.39 MB, less than 80.00%
func restoreArray(adjacentPairs [][]int) []int {
	result := make([]int, len(adjacentPairs)+1)

	hashmap := map[int][]int{}
	for _, v := range adjacentPairs {
		hashmap[v[0]] = append(hashmap[v[0]], v[1])
		hashmap[v[1]] = append(hashmap[v[1]], v[0])
	}

	for k, v := range hashmap {
		if len(v) == 1 {
			result[0] = k
			break
		}
	}

	result[1] = hashmap[result[0]][0]
	for i := 2; i < len(adjacentPairs)+1; i++ {
		adj := hashmap[result[i-1]]
		if result[i-2] == adj[0] {
			result[i] = adj[1]
		} else {
			result[i] = adj[0]
		}
	}

	return result
}
