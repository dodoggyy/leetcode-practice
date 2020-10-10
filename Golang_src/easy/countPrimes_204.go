package easy

// Use sieve of Eratosthenes:
// Time Complexity: O(n*loglogn)
// Space Complexity:O(n)
// Runtime: 12 ms, faster than 64.55%
// Memory Usage: 4.9 MB, less than 5.45%
func countPrimes(n int) int {
	isNotPrime := make([]bool, n)
	result := 0

	for i := 2; i < n; i++ {
		if isNotPrime[i] {
			continue
		} else {
			result++
			for j := 2; j*i < n; j++ {
				isNotPrime[j*i] = true
			}
		}
	}

	return result
}
