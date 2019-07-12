/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class RangeSumQuery_303 {
    // Use DP
    // Time Complexity: O(n) + m*O(1) = O(n+m)
    // Space Complexity:O(n)
    // Runtime: 51 ms, faster than 90.92%
    // Memory Usage: 42.6 MB, less than 47.37%
    class NumArray {
        private int[] mSum;

        public NumArray(int[] nums) {
            if (nums.length == 0) {
                return;
            }
            mSum = new int[nums.length];
            mSum[0] = nums[0];
            for (int i = 1; i < nums.length; i++) {
                mSum[i] = mSum[i - 1] + nums[i];
            }
        }

        public int sumRange(int i, int j) {
            // exception handle
            if ((i > j) || (i < 0) || (j < 0) || (i >= mSum.length) || (j >= mSum.length)) {
                return 0;
            }
            if (i == 0) {
                return mSum[j];
            } else {
                return mSum[j] - mSum[i - 1];
            }
        }
    }
}
