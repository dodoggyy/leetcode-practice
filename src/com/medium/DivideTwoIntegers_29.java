/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class DivideTwoIntegers_29 {
    // 2019年8月9日
    // Use bitwise subtraction
    // Time Complexity: O(logn)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 33.6 MB, less than 6.06%
    public int divide(int dividend, int divisor) {
        if (divisor == 0 || (dividend == Integer.MIN_VALUE && divisor == -1)) {
            return Integer.MAX_VALUE;
        }
        boolean bIsNegitive = (dividend > 0) ^ (divisor > 0); // only one negative number may cause negative 
        long mTmpDividend = Math.abs((long) dividend); // conversion to prevent integer minimum
        long mTmpDivisor = Math.abs((long) divisor);
        int mResult = 0;
        while (mTmpDividend >= mTmpDivisor) {
            int mShift = 0;
            while (mTmpDividend >= mTmpDivisor << mShift) {
                mShift++;
            }
            mResult += 1 << (mShift - 1); // -1 because original shift may bigger than dividend
            mTmpDividend -= mTmpDivisor << (mShift - 1);
        }
        return bIsNegitive ? -mResult : mResult;
    }
}
