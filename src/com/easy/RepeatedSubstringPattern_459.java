/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class RepeatedSubstringPattern_459 {
    // 2019年7月21日
    // Use iterative modulation
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 19 ms, faster than 58.90%
    // Memory Usage: 37.2 MB, less than 100.00%
    public boolean repeatedSubstringPattern(String s) {
        int mLength = s.length();
        for (int i = mLength / 2; i >= 1; i--) {
            if (mLength % i == 0) {
                int mSlice = mLength / i;
                String mSubString = s.substring(0, i);
                StringBuilder mStrBuilder = new StringBuilder();
                for (int j = 0; j < mSlice; j++) {
                    mStrBuilder.append(mSubString);
                }
                if (mStrBuilder.toString().equals(s)) {
                    return true;
                }
            }
        }
        return false;
    }

    // 2019年7月21日
    // Use iterative modulation
    // Time Complexity: O()
    // Space Complexity:O()
    // Runtime: 115 ms, faster than 15.28%
    // Memory Usage: 41.7 MB, less than 5.00%
    public boolean repeatedSubstringPattern2(String s) {
        if (s == null || s.length() == 0) {
            return false;
        }
        int mLength = s.length();

        for (int i = 1; i <= mLength / 2; i++) {
            if ((mLength % i) == 0) {
                String mSubString = s.substring(0, i);
                String[] mStrArray = s.split(mSubString);
                if (mStrArray.length == 0) {
                    return true;
                }
            }
        }
        return false;
    }
}
