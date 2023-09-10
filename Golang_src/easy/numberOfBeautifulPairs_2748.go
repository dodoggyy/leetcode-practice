package easy

// Use math:
// n: nums length
// U: max(nums)
// Time Complexity: O(n*logU))
// Space Complexity:O(k)
// Runtime: 15 ms, faster than 96.30%
// Memory Usage: 5.04 MB, less than 74.07%
func countBeautifulPairs(nums []int) int {
	result := 0

	var gcd func(a, b int) int
	gcd = func(a, b int) int {
		if b == 0 {
			return a
		}
		return gcd(b, a%b)
	}

	cnt := [10]int{}
	for _, v := range nums {
		for i := 1; i < 10; i++ {
			if cnt[i] > 0 && gcd(v%10, i) == 1 {
				result += cnt[i]
			}
		}
		for v >= 10 {
			v /= 10
		}
		cnt[v]++
	}

	return result
}
