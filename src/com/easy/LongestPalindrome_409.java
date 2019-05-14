/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class LongestPalindrome_409 {

    // Array count
    // Time Complexity: O(2n) = O(n)
    // Space Complexity:O(n)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 33.8 MB, less than 100.00%
    public int longestPalindrome(String s) {
        int mResult = 0;
        boolean bContainOdd = false;
        int[] CharsNum = new int[52];
        for (char c : s.toCharArray()) {
            if (c - 'a' < 0) { // Upper case
                CharsNum[c - 'A' + 26]++;
            } else { // Lower case
                CharsNum[c - 'a']++;
            }
        }
        for (int mCount : CharsNum) {
            if (mCount >= 1) {
                if (!bContainOdd && (mCount % 2 == 1)) {
                    bContainOdd = true;
                    mResult += mCount;
                } else {
                    mResult += ((mCount >> 1) << 1);
                    // mResult += (mCount - mCount%2);
                    // mResult += ((mCount/2)*2);
                }

            }

        }

        return mResult;
    }
}
