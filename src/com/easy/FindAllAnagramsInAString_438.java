/**
 * 
 */
package com.easy;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class FindAllAnagramsInAString_438 {
    // 2019年7月19日
    // Use array count
    // Time Complexity: O(n^2)
    // Space Complexity:O(n)
    // Runtime: 487 ms, faster than 15.27%
    // Memory Usage: 38.3 MB, less than 99.74%
    public List<Integer> findAnagrams(String s, String p) {
        List<Integer> mResult = new ArrayList<>();
        if (s == null || p == null || s.length() == 0 || p.length() == 0 || p.length() > s.length()) {
            return mResult;
        }

        int[] mCount = new int[26];
        boolean bIsAnagram = true;
        for (int i = 0; i < p.length(); i++) {
            mCount[p.charAt(i) - 'a']++;
        }
        for (int i = 0; i < (s.length() - p.length() + 1); i++) {
            int[] mTmp = mCount.clone();
            bIsAnagram = true;
            for (int j = 0; j < p.length(); j++) {
                if (--mTmp[s.charAt(i + j) - 'a'] < 0) {
                    bIsAnagram = false;
                    break;
                }

            }
            if (bIsAnagram) {
                mResult.add(i);
            }

        }

        return mResult;
    }

    // 2019年7月19日
    // Use sliding window
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 8 ms, faster than 75.28%
    // Memory Usage: 38.6 MB, less than 98.88%
    public List<Integer> findAnagrams2(String s, String p) {
        List<Integer> mResult = new ArrayList<>();
        if (s == null || p == null || s.length() == 0 || p.length() == 0 || p.length() > s.length()) {
            return mResult;
        }

        int[] mCountP = new int[26];
        int mLengthP = p.length();
        for (int i = 0; i < p.length(); i++) {
            mCountP[p.charAt(i) - 'a']++;
        }
        int[] mCountS = new int[26];
        for (int i = 0; i < s.length(); i++) {
            mCountS[s.charAt(i) - 'a']++;

            if (i >= mLengthP) {
                mCountS[s.charAt(i - mLengthP) - 'a']--;

            }
            if (Arrays.equals(mCountP, mCountS)) {
                mResult.add(i - mLengthP + 1);
            }
        }

        return mResult;
    }
}
