/**
 * 
 */
package com.medium;

import java.util.HashMap;
import java.util.Map;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class DecodeWays_91 {
    public static void main(String[] args) {
        String mStr = "12";
        DecodeWays_91 mTest = new DecodeWays_91();
        System.out.println(mTest.numDecodings2(mStr));
    }

    // Recursion with memorization
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 2 ms, faster than 51.89%
    // Memory Usage: 35.1 MB, less than 99.97%
    private Map<Integer, Integer> mMap = new HashMap<>();

    public int numDecodings(String s) {
        if (s == null || s.length() == 0) {
            return 0;
        }

        return ways(s, 0, s.length() - 1);
    }

    private int ways(String s, int mLeftIndex, int mRightIndex) {
        if (mMap.containsKey(mLeftIndex)) {
            return mMap.get(mLeftIndex);
        }
        // Check out of bound AND judge illegal for 0 at first digit
        if ((mLeftIndex < s.length()) && s.charAt(mLeftIndex) == '0') {
            return 0;
        }
        // Single digit or empty case
        if (mLeftIndex >= mRightIndex) {
            return 1;
        }

        int w = ways(s, mLeftIndex + 1, mRightIndex);
        int mPrefix = (s.charAt(mLeftIndex) - '0') * 10 + (s.charAt(mLeftIndex + 1) - '0');

        if (mPrefix <= 26) {
            w += ways(s, mLeftIndex + 2, mRightIndex);
        }
        mMap.put(mLeftIndex, w);
        return w;
    }

    // Use DP
    // Time Complexity: O(n)
    // Space Complexity: O(1)
    // Runtime: 1 ms, faster than 99.14%
    // Memory Usage: 34.7 MB, less than 100.00%
    public int numDecodings2(String s) {
        if (s == null || s.length() == 0 || s.charAt(0) == '0') {
            return 0;
        }
        if (s.length() == 1) {
            return 1;
        }
        int mLength = s.length();
        int w1 = 1, w2 = 1;
        // consider 4 case:
        // case 1: 0, if s[i] AND s[i-1]s[i] are invalid
        // case 2: dp[i-1]+dp[i-2], if s[i-1] AND s[i-1]s[i] are valid
        // case 3: dp[i-1], if s[i] is valid
        // case 4: dp[i-2], if s[i-1]s[i] is valid
        for (int i = 1; i < mLength; i++) {
            int w = 0;
            if (!bIsValid(s.charAt(i)) && !bIsValid(s.charAt(i - 1), s.charAt(i))) {
                return 0;
            }
            if (bIsValid(s.charAt(i))) {
                w = w1;
            }
            if (bIsValid(s.charAt(i - 1), s.charAt(i))) {
                w += w2;
            }
            w2 = w1;
            w1 = w;
        }

        return w1;
    }

    private boolean bIsValid(char c) {
        return c != '0';
    }

    private boolean bIsValid(char c1, char c2) {
        int mNum = (c1 - '0') * 10 + (c2 - '0');
        return (mNum >= 10) && (mNum <= 26);
    }
}
