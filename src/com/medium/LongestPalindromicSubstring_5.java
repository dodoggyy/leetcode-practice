/**
 * 
 */
package com.medium;


/**
 * @author Chris Lin
 * @version 1.0
 */
public class LongestPalindromicSubstring_5 {
    // 2019年8月6日
    // Use valid palindrome method
    // Time Complexity: O(n^2)
    // Space Complexity:O(1)
    // Runtime: 5 ms, faster than 96.59%
    // Memory Usage: 36.2 MB, less than 100.00%
    private int mMaxLength = 0;
    private int mIndexStart = 0;

    public String longestPalindrome(String s) {
        if (s.length() < 2) {
            return s;
        }
        int mLength = s.length();

        for (int i = 0; i < mLength - 1; i++) {
            searchPalindrome(s, i, i);
            searchPalindrome(s, i, i + 1);
        }
        return s.substring(mIndexStart, mIndexStart + mMaxLength);
    }

    private void searchPalindrome(String s, int mIndexLeft, int mIndexRight) {
        while (mIndexLeft >= 0 && mIndexRight < s.length()) {
            if (s.charAt(mIndexLeft) != s.charAt(mIndexRight)) {
                break;
            }
            mIndexLeft--;
            mIndexRight++;
        }
        if (mMaxLength < mIndexRight - mIndexLeft - 1) {
            mMaxLength = mIndexRight - mIndexLeft - 1;
            mIndexStart = mIndexLeft + 1;
        }
    }

    // 2019年8月6日
    // Use DP
    // Time Complexity: O(n^2)
    // Space Complexity:O(n^2)
    // Runtime: 32 ms, faster than 45.97%
    // Memory Usage: 38.1 MB, less than 43.41%
    public String longestPalindrome2(String s) {
        // if s[i] == s[j] , mDP[i][j] = true
        if (s == null || s.length() == 0) {
            return s;
        }
        int mLength = s.length();
        int mIndexStart = 0, mMaxLength = 0;
        // mDP[i][j] indicates whether substring s starting at index i and
        // ending at j is palindrome
        boolean[][] mDP = new boolean[mLength][mLength];

        for (int i = mLength - 1; i >= 0; i--) {
            for (int j = i; j < mLength; j++) {
                mDP[i][j] = s.charAt(i) == s.charAt(j) && (j - i < 3 || mDP[i + 1][j - 1]);

                if (mDP[i][j] && (j - i + 1 > mMaxLength)) {
                    mIndexStart = i;
                    mMaxLength = j - i + 1;
                }
            }
        }
        return s.substring(mIndexStart, mIndexStart + mMaxLength);
    }
}
