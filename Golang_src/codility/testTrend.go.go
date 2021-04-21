package codility

import "sort"

// Use iterative:
// Time Complexity: O(n)
// Space Complexity:O(1)
func SolutionTrend1(S string) int {
	// write your code in Go 1.4
	result, total := 0, 0

	for _, v := range S {
		total += int(v - '0')
	}

	if total%3 == 0 {
		result++
	}

	for _, v := range S {
		curNum := int(v - '0')
		for i := 0; i < 10; i++ {
			if (i != curNum) && ((total-curNum+i)%3) == 0 {
				result++
			}
		}
	}

	return result
}

// Use sorting:
// Time Complexity: O(K * N^K)
// Space Complexity:O(K)
func SolutionTrend2(A []int, K int) int {
	// write your code in Go 1.4
	size := len(A)
	result := 0
	if size < K {
		return -1
	}

	even, odd := []int{}, []int{}

	for _, v := range A {
		if (v % 2) == 1 {
			odd = append(odd, v)
		} else {
			even = append(even, v)
		}
	}

	sort.Ints(odd)
	sort.Ints(even)

	idxEven, idxOdd := len(even)-1, len(odd)-1

	for K > 0 {
		if (K % 2) == 1 {
			if idxEven >= 0 {
				result += even[idxEven]
				idxEven--
			} else {
				return -1
			}
			K--
		} else if idxEven >= 1 && idxOdd >= 1 {
			if even[idxEven]+even[idxEven-1] <= odd[idxOdd]+odd[idxOdd-1] {
				result += odd[idxOdd] + odd[idxOdd-1]

				idxOdd -= 2
			} else {
				result += even[idxEven] + even[idxEven-1]

				idxEven -= 2
			}
			K -= 2
		} else if idxEven >= 1 {
			result += even[idxEven] + even[idxEven-1]
			idxEven -= 2
			K -= 2
		} else if idxOdd >= 1 {
			result += odd[idxOdd] + odd[idxOdd-1]
			idxOdd -= 2
			K -= 2
		}

	}

	return result
}

// Use math:
// Time Complexity: O(n)
// Space Complexity:O(1)
func SolutionTrend3(A []int, F int, M int) []int {
	// write your code in Go 1.4
	if F < 0 || len(A) == 0 || M < 1 || M > 6 {
		return []int{0}
	}

	result := make([]int, F)
	rollTotal := len(A) + F
	sum := 0

	for _, v := range A {
		sum += v
	}

	remain := M*rollTotal - sum

	if F > remain || remain/F > 6 {
		return []int{0}
	}

	for i := range result {
		result[i] = remain / F
		F--
		remain -= result[i]
	}

	return result
}
