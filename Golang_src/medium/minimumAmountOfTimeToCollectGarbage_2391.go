package medium

// Use prefix sum:
// n: length of garbage collection
// m: average length of garbage
// Time Complexity: O(n*m)
// Space Complexity:O(1)
// Runtime: 129 ms, faster than 95.24%
// Memory Usage: 13.84 MB, less than 71.43%
func garbageCollection(garbage []string, travel []int) int {
	tmpM, tmpP, tmpG := 0, 0, 0

	for i := range travel {
		if i > 0 {
			travel[i] += travel[i-1]
		}
	}

	hasG, hasP, hasM := false, false, false
	for i := len(garbage) - 1; i >= 0; i-- {
		for _, ch := range garbage[i] {
			switch ch {
			case 'G':
				if i > 0 && !hasG {
					hasG = true
					tmpG += travel[i-1]
				}
				tmpG++
			case 'P':
				if i > 0 && !hasP {
					hasP = true
					tmpP += travel[i-1]
				}
				tmpP++
			case 'M':
				if i > 0 && !hasM {
					hasM = true
					tmpM += travel[i-1]
				}
				tmpM++
			}
		}
	}

	return tmpM + tmpP + tmpG
}
