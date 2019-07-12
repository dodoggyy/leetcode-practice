/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class NumberOf1Bits_191 {
    // 2019年7月11日
    // Use Lib bitcount function
    // Time Complexity: O(1)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 27.52%
    // Memory Usage: 33.5 MB, less than 5.00%
    public int hammingWeight(int n) {
        return Integer.bitCount(n);
    }

    // 2019年7月11日
    // Use Bit Manipulation Trick
    // Time Complexity: O(1)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 27.52%
    // Memory Usage: 33.5 MB, less than 5.00%
    public int hammingWeight2(int n) {
        int mSum = 0;
        while (n != 0) {
            mSum++;
            n &= (n - 1);
        }
        return mSum;
    }
}
