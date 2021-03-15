package codility

// Biggest two-digit
// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
func Solution(S string) int {
	// write your code in Go 1.4
	result := 0

	for i := 0; i < len(S)-1; i++ {
		digit1, digit0 := int(S[i]-'0'), int(S[i+1]-'0')
		if digit1 == 0 {
			continue
		}
		result = max(result, digit1*10+digit0)

	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Longest Semi Alternating Substring
// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
func Solution2(S string) int {
	// write your code in Go 1.4
	result, cnt := 0, 0

	var pre byte = '#'

	for i, j := 0, 0; i < len(S); i++ {
		if i == 0 || S[i] != pre {
			pre = S[i]
			cnt = 1
		} else {
			cnt++
			if cnt >= 3 {
				j = i - 1
			}
		}

		result = max(result, i-j+1)
	}

	return result
}

type Tree struct {
	X int
	L *Tree
	R *Tree
}

// Count root to leaf paths having exactly K distinct nodes in a Binary Tree
// Use iterative:
// Time Complexity: O(n*h^2)
// Space Complexity:O(n)
func Solution3(T *Tree) int {
	// write your code in Go 1.4

	if T == nil {
		return 0
	}

	hashmap := map[int]int{}

	var longestPath func(node *Tree) int
	longestPath = func(node *Tree) int {
		if node == nil {
			return len(hashmap)
		}
		if val, exist := hashmap[node.X]; exist {
			hashmap[node.X] = val + 1
		} else {
			hashmap[node.X] = 1
		}

		maxPath := max(longestPath(node.L), longestPath(node.R))

		if _, exist := hashmap[node.X]; exist {
			delete(hashmap, node.X)
		}

		if hashmap[node.X] == 0 {
			delete(hashmap, node.X)
		}

		return maxPath
	}

	return longestPath(T)
}
