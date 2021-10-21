package medium

import (
	"math/rand"
	"time"
)

func rand7() int {
	rand.Seed(time.Now().UnixNano())
	min, max := 1, 7
	return rand.Intn((max - min + 1) + min)
}

// Use Rejection Sampling:
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 19 ms, faster than 20.69%
// Memory Usage: 5.8 MB, less than 37.93%
func rand10() int {
	num := (rand7()-1)*7 + rand7()
	for num > 40 {
		num = (rand7()-1)*7 + rand7()
	}

	return num%10 + 1
}

// Use optimize Rejection Sampling :
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 15 ms, faster than 27.59%
// Memory Usage: 6 MB, less than 13.79%
func rand102() int {
	for {
		num := (rand7()-1)*7 + rand7()
		if num <= 40 {
			return 1 + num%10
		}

		num = (num-40-1)*7 + rand7()
		if num <= 60 {
			return 1 + num%10
		}

		num = (num-60-1)*7 + rand7()
		if num <= 20 {
			return 1 + num%10
		}
	}

}
