package medium

// Use hashmap:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 18 ms, faster than 70.97%
// Memory Usage: 7.06 MB, less than 38.71%
func minDeletions(s string) int {
	result := 0
	cnt := [26]int{}

	for _, v := range s {
		cnt[v-'a']++
	}

	hashmap := map[int]struct{}{}
	
	cur := 0
	for cur < len(cnt) {
		if cnt[cur] == 0 {
			cur++
			continue
		}
		if _, ok := hashmap[cnt[cur]]; !ok {
			hashmap[cnt[cur]] = struct{}{}
			cur++
		} else {
			result++
			cnt[cur]--
		}
	}

	return result
}

