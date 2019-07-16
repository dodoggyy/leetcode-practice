/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ValidPerfectSquare_367 {
    // 2019年7月16日
    // Use built-in library function
    // Time Complexity: O(1)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 33.1 MB, less than 5.07%
    public boolean isPerfectSquare(int num) {
        return Math.sqrt(num) % 1 == 0;
    }

    // 2019年7月16日
    // Use iterative times
    // Time Complexity: O(logn)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 30.35%
    // Memory Usage: 33.1 MB, less than 5.07%
    public boolean isPerfectSquare2(int num) {
        for (int i = 1; i <= num / i; i++) {
            if (i * i == num) {
                return true;
            }
        }
        return false;
    }

    // 2019年7月16日
    // Use binary search
    // Time Complexity: O(logn)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00% 
    // Memory Usage: 32.7 MB, less than 5.07%
    public boolean isPerfectSquare3(int num) {
        long mLow = 0, mHigh = num;

        while (mLow <= mHigh) {
            long mMid = mLow + (mHigh - mLow) / 2;
            long mTmp = mMid * mMid;
            if (mTmp == num) {
                return true;
            } else if (mTmp > num) {
                mHigh = mMid - 1;
            } else { // mTmp < num
                mLow = mMid + 1;
            }
        }
        return false;
    }
}
