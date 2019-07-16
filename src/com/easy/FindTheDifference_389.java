/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class FindTheDifference_389 {
    // 2019年7月17日
    // Use Array count
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 2 ms, faster than 45.17%
    // Memory Usage: 37.7 MB, less than 7.13%
    public char findTheDifference(String s, String t) {
        int[] mCount = new int[26];

        for (int i = 0; i < s.length(); i++) {
            mCount[s.charAt(i) - 'a']++;
        }
        for (int i = 0; i < t.length(); i++) {
            if (--mCount[t.charAt(i) - 'a'] < 0) {
                return t.charAt(i);
            }
        }
        return 0;
    }
}
