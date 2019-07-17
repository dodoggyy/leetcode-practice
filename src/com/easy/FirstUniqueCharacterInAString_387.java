/**
 * 
 */
package com.easy;

import java.util.HashMap;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class FirstUniqueCharacterInAString_387 {
    // 2019年7月17日
    // Use HashMap
    // Time Complexity: O()
    // Space Complexity:O()
    // Runtime: 39 ms, faster than 36.21%
    // Memory Usage: 37.2 MB, less than 99.32%
    public int firstUniqChar(String s) {
        HashMap<Character, Integer> mMap = new HashMap<>();
        for (int i = 0; i < s.length(); i++) {
            mMap.put(s.charAt(i), mMap.getOrDefault(s.charAt(i), 0) + 1);
        }
        for (int i = 0; i < s.length(); i++) {
            if (mMap.get(s.charAt(i)) == 1) {
                return i;
            }
        }
        return -1;
    }

    // 2019年7月17日
    // Use Array count
    // Time Complexity: O(2n) = O(n)
    // Space Complexity:O(1)
    // Runtime: 8 ms, faster than 83.77%
    // Memory Usage: 37.1 MB, less than 99.35%
    public int firstUniqChar2(String s) {
        if (s == null || s.length() == 0) {
            return -1;
        }
        int[] mCount = new int[26];
        for (int i = 0; i < s.length(); i++) {
            mCount[s.charAt(i) - 'a']++;
        }
        for (int i = 0; i < s.length(); i++) {
            if (mCount[s.charAt(i) - 'a'] == 1) {
                return i;
            }
        }
        return -1;
    }
}
