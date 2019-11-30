/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class LongestCommonPrefix_14 {
    // 2019年12月1日
    // Use horizontal scanning
    // Time Complexity: O(S)
    // Space Complexity:O(1)
    // S is the sum of all characters in all strings
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 37.9 MB, less than 80.70%
    public String longestCommonPrefix(String[] strs) {
        if (strs.length == 0) {
            return "";
        }
        StringBuilder mResult = new StringBuilder();
        mResult.append(strs[0]);
        for (int i = 1; i < strs.length; i++) {
            while (strs[i].indexOf(mResult.toString()) != 0) {
                mResult.setLength(mResult.length() - 1);
                if (mResult.length() == 0) {
                    return "";
                }
            }
        }

        return mResult.toString();
    }

    // 2019年12月1日
    // Use vertical scanning
    // Time Complexity: O(S)
    // Space Complexity:O(1)
    // S is the sum of all characters in all strings
    // Runtime: 1 ms, faster than 74.43%
    // Memory Usage: 35.9 MB, less than 100.00%
    public String longestCommonPrefix2(String[] strs) {
        if (strs == null || strs.length == 0) {
            return "";
        }
        for (int i = 0; i < strs[0].length(); i++) {
            char c = strs[0].charAt(i);
            for (int j = 1; j < strs.length; j++) {
                if (i == strs[j].length() || strs[j].charAt(i) != c) {
                    return strs[0].substring(0, i);
                }
            }
        }
        return strs[0];
    }
}
