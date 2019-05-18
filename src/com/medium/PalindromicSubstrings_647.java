/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class PalindromicSubstrings_647 {
    private int mCount = 0;
    
    // Use recursive
    // Time Complexity: O(n^2)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 33.4 MB, less than 100.00%
    public int countSubstrings(String s) {
        for(int i = 0; i < s.length(); i++) {
            helper(s, i, i); // odd case, e.g. aba
            helper(s, i, i+1); // even case, e.g. abba 
        }
        return mCount;
    }

    private void helper(String s, int mIndexLeft, int mIndexRight) {
        while(mIndexLeft >=0 && mIndexRight < s.length() && s.charAt(mIndexLeft) == s.charAt(mIndexRight)) {
            mIndexLeft--;
            mIndexRight++;
            mCount++;
        }
    }
}
