/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ReverseStringII_541 {
    // 2019年7月28日
    // Use iterative swap
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 35.9 MB, less than 100.00%
    public String reverseStr(String s, int k) {
        char[] mCharStr = s.toCharArray();
        int mLength = mCharStr.length;

        for (int i = 0; i < mCharStr.length; i += 2 * k) {
            if (i + k < mLength) {
                reverse(mCharStr, i, i + k - 1);
            } else {
                reverse(mCharStr, i, mLength - 1);
            }
        }

        return new String(mCharStr);
    }

    private void reverse(char[] mCharStr, int mStart, int mEnd) {
        while (mStart < mEnd) {
            char mTmp = mCharStr[mStart];
            mCharStr[mStart] = mCharStr[mEnd];
            mCharStr[mEnd] = mTmp;
            mStart++;
            mEnd--;
        }
    }
}
