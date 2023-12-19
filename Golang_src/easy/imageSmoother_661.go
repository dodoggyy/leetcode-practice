package easy

// Use Bit Manipulation:
// Time Complexity: O(m*n)
// Space Complexity:O(1)
// Runtime: 33 ms, faster than 77.78%
// Memory Usage: 7.14 MB, less than 88.89%
func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum, cnt := 0, 0

			for x := i - 1; x <= i+1; x++ {
				for y := j - 1; y <= j+1; y++ {
					if x >= 0 && x < m && y >= 0 && y < n {
						sum += img[x][y] & 255
						cnt++
					}
				}
			}

			img[i][j] |= (sum / cnt) << 8
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			img[i][j] >>= 8
		}
	}

	return img
}
