package easy

// Use greddy algorithm:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 16 ms, faster than 87.59%
// Memory Usage: 5.9 MB, less than 94.99%
func canPlaceFlowers(flowerbed []int, n int) bool {
	index, count := 0, 0
	for index < len(flowerbed) {
		if flowerbed[index] == 0 && (index == 0 || flowerbed[index-1] == 0) && (index == len(flowerbed)-1 || flowerbed[index+1] == 0) {
			flowerbed[index] = 1
			index++
			count++
		}
		if count >= n {
			return true
		}
		index++
	}
	return false
}
