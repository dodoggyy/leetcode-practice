package medium

// Use BFS:
// C: 4 (ACGT)
// n: gene sequence length
// m: bank length
// Time Complexity: O(C*n*m)
// Space Complexity:O(n*m)
// Runtime: 1 ms, faster than 86.41%
// Memory Usage: 1.99 MB, less than 55.34%
func minMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}
	hashset := map[string]bool{}
	for _, v := range bank {
		hashset[v] = true
	}
	if _, ok := hashset[endGene]; !ok {
		return -1
	}

	queue := []string{startGene}
	step := 0
	for len(queue) > 0 {
		step++
		tmp := queue
		queue = []string{}
		for _, cur := range tmp {
			for i, x := range cur {
				for _, y := range "ACGT" {
					if y != x {
						next := cur[:i] + string(y) + cur[i+1:]
						if _, ok := hashset[next]; ok {
							if next == endGene {
								return step
							}
							delete(hashset, next)
							queue = append(queue, next)
						}
					}
				}
			}
		}
	}
	return -1
}
