/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.1
 */
public class MissingNumber_268 {

    // simple calculation
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 41 MB, less than 61.27%
    public int missingNumber(int[] nums) {
        int mResult = 0;
        mResult = (nums.length + 1) * nums.length / 2;

        for (int num : nums) {
            mResult -= num;
        }

        return mResult;
    }

    // XOR operation
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 41.3 MB, less than 59.68%
    public int missingNumber2(int[] nums) {
        int mResult = 0;
        for (int i = 0; i < nums.length; i++) {
            mResult = mResult ^ i ^ nums[i];
        }
        // need XOR array biggest number ex: [0,1] so
        // length is 2
        mResult ^= nums.length;

        return mResult;
    }
}
