/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class GuessNumberHigherOrLower_374 {
    public int guess(int num) {
        return 0;
    }

    // 2019年7月17日
    // Use binary search
    // Time Complexity: O(logn)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 32.9 MB, less than 100.00%
    public int guessNumber(int n) {
        int mLeft = 1, mRight = n;
        int mResult = 0;

        while (mLeft <= mRight) {
            int mMid = mLeft + (mRight - mLeft) / 2;
            if (guess(mMid) == -1) {
                mRight = mMid - 1;
            } else if (guess(mMid) == 1) {
                mLeft = mMid + 1;
            } else {
                mResult = mMid;
                break;
            }

        }

        return mResult;
    }
}
