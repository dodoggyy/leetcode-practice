/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class MaxConsecutiveOnes_485 {
    // 2019年7月23日
    // Use iterative
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 2 ms, faster than 99.67%
    // Memory Usage: 38.2 MB, less than 100.00%
    public int findMaxConsecutiveOnes(int[] nums) {
        int mResult = 0;
        int mTmp = 0;

        for (int i = 0; i < nums.length; i++) {
            if (nums[i] == 1) {
                mTmp++;
            } else { // 0 case
                mResult = Math.max(mTmp, mResult);
                mTmp = 0;
            }
        }
        mResult = Math.max(mTmp, mResult);

        return mResult;
    }

}
