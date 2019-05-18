/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class IsomorphicStrings_205 {

    // Array count
    // Time Complexity: O(n)
    // Space Complexity: O(2n) = O(n)
    // Runtime: 4 ms, faster than 87.99%
    // Memory Usage: 34.9 MB, less than 87.92%
    public boolean isIsomorphic(String s, String t) {
        int[] mFirstStrCount = new int[256];
        int[] mSecondStrCount = new int[256];
        int mLength = s.length();

        for (int i = 0; i < mLength; i++) {
            mFirstStrCount[s.charAt(i)] += (i + 1);
            mSecondStrCount[t.charAt(i)] += (i + 1);
            if (mFirstStrCount[s.charAt(i)] != mSecondStrCount[t.charAt(i)]) {
                return false;
            }
        }

        return true;

    }
}
