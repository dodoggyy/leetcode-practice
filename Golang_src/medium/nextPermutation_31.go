package medium

// Use 2-round iterative scan:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 9 ms, faster than 6.16%
// Memory Usage: 2.8 MB, less than 13.77%
func nextPermutation(nums []int) {
	size := len(nums)
	if size <= 1 {
		return
	}

	i, j := size-2, size-1
	k := size - 1

	// find nums[i] < nums[j]
	// 從後往前找到第一個【相鄰升序對】，即nums[i]<nums[i+1]。此時nums[i+1,end)為降序。
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	if i >= 0 {
		// find nums[i] < nums[k]
		// 在區間[i+1,end)中，從後往前找到第一個大於nums[i]的元素A[k]
		for nums[i] >= nums[k] {
			k--
		}
		// swap nums[i], nums[k]
		// 交換nums[i]和nums[k]，此時nums[i+1,end)一定還是降序，因為nums[j]是從右側起第一個大於nums[i]的值
		nums[i], nums[k] = nums[k], nums[i]
	}

	reverse := func(arr []int) {
		size := len(arr)
		for i := 0; i < size/2; i++ {
			arr[i], arr[size-1-i] = arr[size-1-i], arr[i]
		}
	}

	// reverse nums[j:]
	// 反轉 nums[i+1,end)，變成升序
	reverse(nums[j:])
}
