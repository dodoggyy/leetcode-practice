/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class AddStrings_415 {
    // 2019年7月18日
    // Use sequential parsing
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 4 ms, faster than 38.26%
    // Memory Usage: 36.3 MB, less than 99.80%
    public String addStrings(String num1, String num2) {
        // same concept as LeetCode 67. Add Binary
        StringBuilder mBuilder = new StringBuilder();
        int mLength1 = num1.length();
        int mLength2 = num2.length();
        int mLength = Math.max(mLength1, mLength2);
        int mCarry = 0;
        for (int i = 0; i < mLength; i++) {
            int mTmp1 = (mLength1 > i) ? (num1.charAt(mLength1 - i - 1) - '0') : 0;
            int mTmp2 = (mLength2 > i) ? (num2.charAt(mLength2 - i - 1) - '0') : 0;
            mBuilder.insert(0, (mTmp1 + mTmp2 + mCarry) % 10);
            mCarry = (mTmp1 + mTmp2 + mCarry) / 10;
        }
        if (mCarry != 0) {
            mBuilder.insert(0, mCarry);
        }

        return mBuilder.toString();
    }
}
