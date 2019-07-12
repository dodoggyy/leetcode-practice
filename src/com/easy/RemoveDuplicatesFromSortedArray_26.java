/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class RemoveDuplicatesFromSortedArray_26 {
    // Use index count
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 38.8 MB, less than 99.99%
    public int removeDuplicates(int[] nums) {
        if (nums.length == 0) {
            return 0;
        }
        int mIndex = 0;
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] != nums[i - 1]) {
                nums[++mIndex] = nums[i];
            }
        }

        return (mIndex + 1);
    }
}
