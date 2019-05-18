/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class CountBinarySubstrings_696 {
    public static void main(String[] args) {
        String s = "000111000";//"00110011";

        System.out.println(countBinarySubstrings(s));
    }

    // Two pointer judgement
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 13 ms, faster than 72.15%
    // Memory Usage: 36.3 MB, less than 99.71%
    public static int countBinarySubstrings(String s) {
        int mPreviousLength = 0, mCurrentLength = 1, mResult = 0;
            
        for (int i = 1; i < s.length(); i++) {
            if(s.charAt(i) == s.charAt(i - 1)) {
                mCurrentLength++;
            } else {
                mPreviousLength = mCurrentLength;
                mCurrentLength = 1;
            }
            if(mPreviousLength >= mCurrentLength) {
                mResult++;
            }
        }
        return mResult;
    }

}
