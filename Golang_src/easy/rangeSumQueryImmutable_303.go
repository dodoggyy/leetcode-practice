package easy

// Use prefix sum:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 28 ms, faster than 88.86%
// Memory Usage: 9.6 MB, less than 65.94%
type NumArray struct {
	arr []int
}

func ConstructorNumArray(nums []int) NumArray {
	arr := make([]int, len(nums))

	tmp := 0
	for i := range nums {
		tmp += nums[i]
		arr[i] = tmp
	}

	return NumArray{arr: arr}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.arr[right]
	}
	return this.arr[right] - this.arr[left-1]
}
