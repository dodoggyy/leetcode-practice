/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class NumberOfSegmentsInAString_434 {
    // 2019年7月19日
    // Use built-in library
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 2 ms, faster than 36.20%
    // Memory Usage: 34.2 MB, less than 99.31%
    public int countSegments(String s) {
        if (s == null || s.trim().length() == 0) {
            return 0;
        }

        return s.trim().split("\\s+").length;
    }

    // 2019年7月19日
    // Use in-place
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 33.9 MB, less than 99.31%
    public int countSegments2(String s) {
        int mCount = 0;

        for (int i = 0; i < s.length(); i++) {
            if ((i == 0 || s.charAt(i - 1) == ' ') && s.charAt(i) != ' ') {
                mCount++;
            }
        }

        return mCount;
    }
}
