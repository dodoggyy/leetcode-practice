/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.1
 */
public class BinaryNumberwithAlternatingBits_693 {
    public static void main(String[] args) {
        int n = 5;
        hasAlternatingBits(n);
    }

    // bitwise operation 1
    // Time Complexity: O(1)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 32 MB, less than 100.00%
    public static boolean hasAlternatingBits(int n) {
        // System.out.println(Integer.toBinaryString(n).length());
        return ((n > 0) && ((n & (n >> 1)) == 0)
                && (Integer.bitCount((n | (n >> 1))) == Integer.toBinaryString(n).length()));
    }

    // bitwise operation 2
    // Time Complexity: O(1)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 32 MB, less than 100.00%
    public static boolean hasAlternatingBits2(int n) {
        int mTmp = (n ^ (n >> 1));

        return ((mTmp & (mTmp + 1)) == 0);
    }
}
