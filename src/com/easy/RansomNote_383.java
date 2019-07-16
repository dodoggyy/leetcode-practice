/**
 * 
 */
package com.easy;

import java.util.HashMap;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class RansomNote_383 {
    public static void main(String[] args) {
        RansomNote_383 mTest = new RansomNote_383();
        String ransomNote = "gdcff";
        String magazine = "hiagejhbfdfgbdbccgjfbefgbdfiajejfabefecjcjbgidcdfhbigecbiediadg";

        System.out.println(mTest.canConstruct(ransomNote, magazine));
    }

    // 2019年7月17日
    // Use HashMap
    // Time Complexity: O(n + m)
    // Space Complexity:O(n)
    // m: ransomNote's length, n: magzine's length
    // Runtime: 23 ms, faster than 28.91% 
    // Memory Usage: 38 MB, less than 99.50%
    public boolean canConstruct(String ransomNote, String magazine) {
        int mLengthNote = ransomNote.length();
        int mLengthMagzine = magazine.length();
        if (mLengthMagzine < mLengthNote) {
            return false;
        }
        HashMap<Character, Integer> mMap = new HashMap<>();
        for (int i = 0; i < mLengthMagzine; i++) {
            mMap.put(magazine.charAt(i), mMap.getOrDefault(magazine.charAt(i), 0) + 1);
        }
        for (int i = 0; i < mLengthNote; i++) {
            char mTmpChar = ransomNote.charAt(i);
            if (!mMap.containsKey(mTmpChar) || mMap.get(mTmpChar) <= 0) {
                return false;
            } else {
                mMap.put(mTmpChar, mMap.get(mTmpChar) - 1);
            }
        }
        return true;
    }

    // 2019年7月17日
    // Use Array to count
    // Time Complexity: O(n + m)
    // Space Complexity:O(n)
    // m: ransomNote's length, n: magzine's length
    // Runtime: 3 ms, faster than 76.53%
    // Memory Usage: 38 MB, less than 99.50%
    public boolean canConstruct2(String ransomNote, String magazine) {
        int mLengthNote = ransomNote.length();
        int mLengthMagzine = magazine.length();
        if (mLengthMagzine < mLengthNote) {
            return false;
        }
        int[] mSum = new int[26];
        for (int i = 0; i < mLengthMagzine; i++) {
            mSum[magazine.charAt(i) - 'a']++;
        }
        for (int i = 0; i < mLengthNote; i++) {
            if (--mSum[ransomNote.charAt(i) - 'a'] < 0) {
                return false;
            }
        }
        return true;
    }
}
