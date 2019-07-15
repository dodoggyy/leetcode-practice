/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class FirstBadVersion_278 {
    public boolean isBadVersion(int version) {
        return true;
    }

    // 2019年7月15日
    // Use binary search
    // Time Complexity: O(logn)
    // Space Complexity:O(1)
    // Runtime: 10 ms, faster than 99.62%
    // Memory Usage: 33.1 MB, less than 100.00%
    public int firstBadVersion(int n) {
        int mLow = 1;
        int mHigh = n;

        while (mLow < mHigh) {
            int mMid = mLow + (mHigh - mLow) / 2;
            if (isBadVersion(mMid)) {
                mHigh = mMid;
            } else {
                mLow = mMid + 1;
            }
        }

        return mLow;
    }
}
