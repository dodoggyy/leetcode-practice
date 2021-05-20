package easy

import "math"

// Use array count:
// Time Complexity: O(s + n*(m+s) + s*m) = O(n*m + m) = O(n*m)
// Space Complexity: O(s) = O(1)
// s : total character set = 26
// m: average length of word
// n: total words
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.9 MB, less than 61.54%
func commonChars(words []string) []string {
	result := []string{}
	minFreq := make([]int, 26)
	// O(s)
	for i := range minFreq {
		minFreq[i] = math.MaxInt32
	}

	// total: O(n*(m+s))
	// O(n)
	for _, word := range words {
		freq := make([]int, 26)
		// O(m)
		for _, ch := range word {
			freq[ch-'a']++
		}
		// O(s)
		for i := range minFreq {
			if freq[i] < minFreq[i] {
				minFreq[i] = freq[i]
			}
		}
	}

	//O(s*m)
	for i := range minFreq {
		for j := 0; j < minFreq[i]; j++ {
			result = append(result, string('a'+i))
		}
	}

	return result
}
