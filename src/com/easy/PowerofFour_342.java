/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.1
 */
public class PowerofFour_342 {

    public static void main(String[] args) {
        int num = 16;
        PowerofFour_342 mTest = new PowerofFour_342();
        mTest.isPowerOfFour2(num);
    }

    // bitwise operation
    // Time Complexity: O(1)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 82.30%
    // Memory Usage: 32.4 MB, less than 100.00%
    public boolean isPowerOfFour(int num) {
        // (positive) AND (power of 2) AND (power of 4)
        return ((num > 0) && ((num & (num - 1)) == 0) && ((num & 0x55555555) != 0));
    }

    // mod 4 judgement
    // Time Complexity: O(logn)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 82.30%
    // Memory Usage: 32.4 MB, less than 100.00%
    public boolean isPowerOfFour2(int num) {
        if (num <= 0) {
            return false;
        }

        while (num > 1) {
            if ((num % 4) != 0) {
                return false;
            }
            num /= 4;
            // System.out.println(num);
        }
        return true;
    }

    // Use regular expression
    // Time Complexity: O(logn)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 82.30%
    // Memory Usage: 32.4 MB, less than 100.00%
    public boolean isPowerOfFour3(int num) {
        // if n to the b th
        // then regular expression with base n may be 10*
        // e.g. 1 => 1, 4 => 10, 16 => 100, 64 => 1000
        return Integer.toString(num, 4).matches("10*");
    }
}
