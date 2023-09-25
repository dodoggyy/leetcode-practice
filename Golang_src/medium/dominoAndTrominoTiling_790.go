package medium

// Use dp:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.4 MB, less than 25.76%
func numTilings(n int) int {
	mod := int(1e9 + 7)

	// f[i][j]
	// j = 0 means all empty
	// j = 1 means all full
	// j = 2 means up full & down empty
	// j = 3 means up empty & down full

	f := make([][4]int, n+1)
	f[1][0], f[1][1] = 1, 1
	for i := 2; i <= n; i++ {
		f[i][0] = f[i-1][1]
		cal := 0
		for j := 0; j < 4; j++ {
			cal = (cal + f[i-1][j]) % mod
		}
		f[i][1] = cal
		f[i][2] = (f[i-1][0] + f[i-1][3]) % mod
		f[i][3] = (f[i-1][0] + f[i-1][2]) % mod
	}

	return f[n][1]
}

// Use optimize dp:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 1.89 MB, less than 98.48%
func numTilings2(n int) int {
	mod := int(1e9 + 7)

	// f[i][j]
	// j = 0 means all empty
	// j = 1 means all full
	// j = 2 means up full & down empty
	// j = 3 means up empty & down full

	f := [2][4]int{}
	f[1][0], f[1][1] = 1, 1
	for i := 2; i <= n; i++ {
		cur := i & 1
		pre := (i - 1) & 1
		f[cur][0] = f[pre][1]
		cal := 0
		for j := 0; j < 4; j++ {
			cal = (cal + f[pre][j]) % mod
		}
		f[cur][1] = cal
		f[cur][2] = (f[pre][0] + f[pre][3]) % mod
		f[cur][3] = (f[pre][0] + f[pre][2]) % mod
	}

	return f[n&1][1]
}
