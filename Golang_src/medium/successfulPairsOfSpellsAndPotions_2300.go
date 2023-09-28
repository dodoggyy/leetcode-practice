package medium

import "sort"

// Use sort + two pointers:
// Time Complexity: O(mlogm+nlogn)
// Space Complexity:O(m+logn)
// Runtime: 232 ms, faster than 27.78%
// Memory Usage: 15.9 MB, less than 5.56%
func successfulPairs(spells []int, potions []int, success int64) []int {
	m, n := len(spells), len(potions)
	result := make([]int, m)

	sortSpells := make([][2]int, m)
	for i := range spells {
		sortSpells[i][0] = spells[i]
		sortSpells[i][1] = i
	}

	sort.Ints(potions)
	sort.Slice(sortSpells, func(a, b int) bool {
		return sortSpells[a][0] < sortSpells[b][0]
	})

	curIdx := n - 1

	for _, v := range sortSpells {
		spell := v[0]
		idx := v[1]

		for curIdx >= 0 && int64(spell*potions[curIdx]) >= success {
			curIdx--
		}
		result[idx] = n - (curIdx + 1)
	}

	return result
}

// Use sort + binary search:
// Time Complexity: O((m+n)*logn)
// Space Complexity:O(1)
// Runtime: 191 ms, faster than 72.22%
// Memory Usage: 12.51 MB, less than 30.86%
func successfulPairs2(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	m, n := len(spells), len(potions)
	result := make([]int, m)

	for i := range spells {
		l, r := 0, n
		for l < r {
			mid := l + (r-l)/2
			if int64(spells[i])*int64(potions[mid]) < success {
				l = mid + 1
			} else {
				r = mid
			}
		}
		result[i] = n - l
	}

	return result
}
