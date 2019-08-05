/**
 * 
 */
package com.medium;

import java.util.HashMap;
import java.util.HashSet;
import java.util.Set;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class LongestSubstringWithoutRepeatingCharacters_3 {
    // 2019年8月5日
    // Use slide window via HashSet
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 9 ms, faster than 50.07%
    // Memory Usage: 36.2 MB, less than 99.88%
    public int lengthOfLongestSubstring(String s) {
        int mResult = 0, mLength = s.length();
        int mIndexLeft = 0, mIndexRight = 0;
        Set<Character> mSet = new HashSet<>();
        while (mIndexLeft < mLength && mIndexRight < mLength) {
            if (mSet.contains(s.charAt(mIndexRight))) {
                mSet.remove(s.charAt(mIndexLeft++));
            } else {
                mSet.add(s.charAt(mIndexRight++));
                mResult = Math.max(mResult, mIndexRight - mIndexLeft);
            }

        }

        return mResult;
    }

    // 2019年8月5日
    // Use slide window via HashMap
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 10 ms, faster than 40.39%
    // Memory Usage: 37.2 MB, less than 98.22%
    public int lengthOfLongestSubstring2(String s) {
        int mResult = 0, mLength = s.length();
        HashMap<Character, Integer> mMap = new HashMap<>();
        for (int i = 0, j = 0; j < mLength; j++) {
            if (mMap.containsKey(s.charAt(j))) {
                i = Math.max(i, mMap.get(s.charAt(j)));
            }
            mResult = Math.max(mResult, j - i + 1);
            mMap.put(s.charAt(j), j + 1);
        }
        return mResult;
    }

    // 2019年8月5日
    // Use slide window via array (assume ASCI src)
    // Time Complexity: O(n)
    // Space Complexity:O(k)
    // k: length of charset
    // Runtime: 3 ms, faster than 91.56%
    // Memory Usage: 37.9 MB, less than 67.01%
    public int lengthOfLongestSubstring3(String s) {
        int mResult = 0, mLength = s.length();
        int[] mSet = new int[128];
        for (int i = 0, j = 0; j < mLength; j++) {
            i = Math.max(i, mSet[s.charAt(j)]);
            mResult = Math.max(mResult, j - i + 1);
            mSet[s.charAt(j)] = j + 1;
        }
        return mResult;
    }
}
