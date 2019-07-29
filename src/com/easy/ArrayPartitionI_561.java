/**
 * 
 */
package com.easy;

import java.util.Arrays;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ArrayPartitionI_561 {
    // 2019年7月30日
    // Use iterative
    // Time Complexity: O(nlogn)
    // Space Complexity:O(1)
    // Runtime: 15 ms, faster than 44.00%
    // Memory Usage: 39.7 MB, less than 100.00%
    public int arrayPairSum(int[] nums) {
        Arrays.sort(nums);
        int mResult = 0;
        for (int i = 0; i < nums.length; i += 2) {
            mResult += nums[i];
        }
        return mResult;
    }
}
