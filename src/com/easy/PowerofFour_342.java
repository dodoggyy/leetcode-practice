/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class PowerofFour_342 {

    public static void main(String[] args) {
        int num = 16;
        isPowerOfFour2(num);
    }

    // bitwise operation
    // Runtime: 1 ms, faster than 82.30%
    // Memory Usage: 32.4 MB, less than 100.00%
    public boolean isPowerOfFour(int num) {
        // (positive) AND (power of 2) AND (power of 4)
        return ((num > 0) && ((num & (num - 1)) == 0) && ((num & 0x55555555) != 0));
    }

    // mod 4 judgement
    // Runtime: 1 ms, faster than 82.30%
    // Memory Usage: 32.4 MB, less than 100.00%
    public static boolean isPowerOfFour2(int num) {
        if (num <= 0) {
            return false;
        }

        while (num > 1) {
            if ((num % 4) != 0) {
                return false;
            }
            num /= 4;
            //System.out.println(num);
        }
        return true;
    }
}
