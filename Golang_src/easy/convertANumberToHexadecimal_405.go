package easy

// Use iterative:
// Time Complexity: O(logn)
// Space Complexity:O(logn)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 1.9 MB, less than 100.00%
func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	hexList := []byte("0123456789abcdef")
	num &= 0xffffffff

	result := ""

	for num != 0 {
		result = string(hexList[num%16]) + result
		num >>= 4
	}

	return result
}

// Use recursive:
// Time Complexity: O(logn)
// Space Complexity:O(logn)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 1.9 MB, less than 100.00%
func toHex2(num int) string {
	if num < 0 {
		return toHex2(num & 0xffffffff)
	}
	hexList := []byte("0123456789abcdef")
	if num < 16 {
		return string(hexList[num])
	}

	return toHex2(num>>4) + string(hexList[num%16])
}
