/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class FactorialTrailingZeroes_172 {
    // 2019年7月11日
    // Count 5's multiple count
    // Time Complexity: O(1ogn)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00% 
    // Memory Usage: 33.3 MB, less than 5.17%
    public int trailingZeroes(int n) {
        int mResult = 0;
        while (n > 0) {
            n /= 5;
            mResult += n;
            
        }

        return mResult;
    }
}
