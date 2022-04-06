package medium

// Use Counting with Cases:
// Time Complexity: O(n + w^2)
// Space Complexity:O(w)
// n: the length of arr
// w: the maximum possible value of arr[i]
// Runtime: 13 ms, faster than 100.00%
// Memory Usage: 3.2 MB, less than 100.00%
func threeSumMulti(arr []int, target int) int {
	result := 0
	mod := 1000000007
	cnt := make([]int, 101)
	for _, v := range arr {
		cnt[v]++
	}

	cal := func(i, j, k int) int {
		if i == j && j == k {
			num := cnt[i]
			return num * (num - 1) * (num - 2) / 6
		}
		if i == j && j != k {
			return cnt[i] * (cnt[i] - 1) * cnt[k] / 2
		}
		if i != j && j == k {
			return cnt[i] * cnt[j] * (cnt[j] - 1) / 2
		}

		return cnt[i] * cnt[j] * cnt[k]
	}

	for i := range cnt {
		if cnt[i] == 0 {
			continue
		}
		for j := i; j < len(cnt); j++ {
			if cnt[j] == 0 {
				continue
			}
			k := target - i - j
			if k < j {
				break
			}
			if k > 100 || cnt[k] == 0 {
				continue
			}

			result += cal(i, j, k) % mod
		}
	}

	return result
}
