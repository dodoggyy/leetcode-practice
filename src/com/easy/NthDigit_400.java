/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class NthDigit_400 {
    // 2019年7月17日
    // Find Math rule
    // Time Complexity: O(logn)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 33.2 MB, less than 5.18%
    public int findNthDigit(int n) {
        // find rule that
        // 1 to 9 => length: 1, total:9
        // 10 to 99 => length: 2, total:90
        // 100 to 999 => length: 3, total:900
        int mDigitLength = 1;
        long mDigitNum = 9; // prevent overflow

        while (n > mDigitLength * mDigitNum) {
            n -= (int) mDigitLength * mDigitNum;
            mDigitLength++;
            mDigitNum *= 10;
        }

        int mIndexRange = (n - 1) / mDigitLength;
        int mIndexNum = (n - 1) % mDigitLength;

        int mNum = (int) Math.pow(10, mDigitLength - 1) + mIndexRange;

        return Integer.parseInt((mNum + "").charAt(mIndexNum) + "");
    }
}
