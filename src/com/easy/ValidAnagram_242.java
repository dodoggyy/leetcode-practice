/**
 * 
 */
package com.easy;

import java.util.HashMap;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ValidAnagram_242 {
    
    // Use hash to count
    // Time Complexity: O(2n) = O(n)
    // Space Complexity:O(n)
    // Runtime: 23 ms, faster than 11.73%
    // Memory Usage: 35.8 MB, less than 86.27%
    public boolean isAnagram(String s, String t) {
        if (s.length() != t.length()) {
            return false;
        }
        
        HashMap<Character, Integer> mMap = new HashMap<>();
        for(int i = 0; i < s.length(); i++) {
            mMap.put(s.charAt(i), mMap.getOrDefault(s.charAt(i), 0) + 1);
            mMap.put(t.charAt(i), mMap.getOrDefault(t.charAt(i), 0) - 1);
        }
        
        for(Character key:mMap.keySet()) {
            if(mMap.get(key) != 0) {
                return false;
            }
        }
        
        return true;
    }
    
    // Use integer array to count
    // Time Complexity: O(3n) = O(n)
    // Space Complexity:O(n)
    // Runtime: 4 ms, faster than 74.01%
    // Memory Usage: 35.2 MB, less than 99.60%
    public boolean isAnagram2(String s, String t) {
        if (s.length() != t.length()) {
            return false;
        }
        
        int [] mCharsNum = new int[26];
        for (char c : s.toCharArray()) {
            mCharsNum[ c - 'a']++;
        }
        for (char c : t.toCharArray()) {
            mCharsNum[ c - 'a']--;
        }
        for(int mCount: mCharsNum) {
            if(mCount != 0) {
                return false;
            }
        }
        return true;
    }
}
