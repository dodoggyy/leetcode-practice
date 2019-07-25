/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ValidPalindrome_125 {
    public static void main(String[] args) {
        ValidPalindrome_125 mTest = new ValidPalindrome_125();
        String s = "0P";// ".,";//"race a car";
        System.out.println(mTest.isPalindrome(s));
    }

    // 2019年7月25日
    // Use dual index comparison
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 3 ms, faster than 96.39%
    // Memory Usage: 37 MB, less than 99.95%
    public boolean isPalindrome(String s) {
        s = s.toUpperCase();
        int mIndexLeft = 0, mIndexRight = s.length() - 1;
        while (mIndexLeft <= mIndexRight) {
            char mHead = s.charAt(mIndexLeft);
            char mTail = s.charAt(mIndexRight);
            if (!checkAlphabet(mHead)) {
                mIndexLeft++;
            } else if (!checkAlphabet(mTail)) {
                mIndexRight--;
            } else {
                if (mHead != mTail) {
                    return false;
                }
                mIndexLeft++;
                mIndexRight--;
            }
        }
        return true;
    }

    // Can replace with build-in library Character.isLetterOrDigit(c)
    private boolean checkAlphabet(char c) {
        if (c >= 65 && c <= 90) { // Upper Case
            return true;
        } else if (c >= 97 && c <= 122) { // Lower Case
            return true;
        } else if (c >= 48 && c <= 57) { // Digit Case
            return true;
        } else {
            return false;
        }
    }

    // 2019年7月25日
    // Use reverse String comparison
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 16 ms, faster than 26.66%
    // Memory Usage: 38.6 MB, less than 79.91%
    public boolean isPalindrome2(String s) {
        s = s.toLowerCase().replaceAll("\\W+", "");
        return s.equals(new StringBuilder().append(s).reverse().toString());
    }
}
