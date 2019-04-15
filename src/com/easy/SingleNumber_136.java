/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class SingleNumber_136 {
    
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 41.7 MB, less than 20.36%
    public int singleNumber(int[] nums) {
        int mResult = 0;
        for (int num : nums) {
            mResult ^= num;
        }

        return mResult;
    }
}
