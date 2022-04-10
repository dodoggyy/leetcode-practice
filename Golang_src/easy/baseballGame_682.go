package easy

import "strconv"

// Use stack:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 6 ms, faster than 22.22%
// Memory Usage: 2.7 MB, less than 22.22%
func calPoints(ops []string) int {
	result := 0
	tmp := []int{}

	for _, str := range ops {
		switch str {
		case "+":
			if len(tmp) < 2 {
				return 0
			}
			tmp = append(tmp, tmp[len(tmp)-2]+tmp[len(tmp)-1])

		case "D":
			if len(tmp) < 1 {
				return 0
			}
			tmp = append(tmp, tmp[len(tmp)-1]*2)
		case "C":
			if len(tmp) < 1 {
				return 0
			}
			tmp = tmp[:len(tmp)-1]
		default:
			num, err := strconv.Atoi(str)
			if err != nil {
				return 0
			}
			tmp = append(tmp, num)
		}
	}

	for _, v := range tmp {
		result += v
	}

	return result
}

// Use stack optimization:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.7 MB, less than 22.22%
func calPoints2(ops []string) int {
	result := 0
	tmp := []int{}

	for _, str := range ops {
		size := len(tmp)
		switch str {
		case "+":
			tmp = append(tmp, tmp[size-2]+tmp[size-1])
			result += tmp[size]
		case "D":
			tmp = append(tmp, tmp[size-1]*2)
			result += tmp[size]
		case "C":
			result -= tmp[size-1]
			tmp = tmp[:size-1]
		default:
			num, _ := strconv.Atoi(str)
			tmp = append(tmp, num)
			result += tmp[size]
		}
	}

	return result
}
