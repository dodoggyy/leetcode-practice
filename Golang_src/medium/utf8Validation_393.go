package medium

// Use Bit Manipulation:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 28 ms, faster than 16.67%
// Memory Usage: 5.4 MB, less than 20.00%
func validUtf8(data []int) bool {
	mask1, mask2 := 1<<7, 1<<7|1<<6

	getByte := func(num int) int {
		if num&mask1 == 0 {
			return 1
		}
		n := 0
		for i := mask1; num&i != 0; i >>= 1 {
			n++
			if n > 4 {
				return -1
			}
		}
		if n >= 2 {
			return n
		}
		return -1
	}

	for i := 0; i < len(data); {
		n := getByte(data[i])
		if (n < 0) || (i+n > len(data)) {
			return false
		}
		for _, v := range data[i+1 : i+n] {
			if v&mask2 != mask1 {
				return false
			}
		}
		i += n
	}

	return true
}
