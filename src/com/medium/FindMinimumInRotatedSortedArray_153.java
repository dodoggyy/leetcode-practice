/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class FindMinimumInRotatedSortedArray_153 {

    // 2019年10月2日
    // Use linear search
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 37.9 MB, less than 100.00%
    public int findMin(int[] nums) {
        for(int i = 0; i < nums.length - 1;i++) {
            if(nums[i] > nums[i+1]) {
                return nums[i+1];
            }
        }
        return nums[0];
    }

    // 2019年10月2日
    // Use binary search
    // Time Complexity: O(logn)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 38.4 MB, less than 84.09%
    public int findMin2(int[] nums) {
        int mIndexLow = 0, mIndexHigh = nums.length - 1;

        while (mIndexLow < mIndexHigh) {
            int mIndexMid = mIndexLow + (mIndexHigh - mIndexLow) / 2;
            if (nums[mIndexMid] < nums[mIndexHigh]) {
                mIndexHigh = mIndexMid;
            } else {
                mIndexLow = mIndexMid + 1;
            }
        }

        return nums[mIndexHigh];
    }
}
