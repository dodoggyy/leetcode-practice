/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 * 
 */
public class AddDigits_258 {
    // Use recursive
    // Time Complexity: O(logn)
    // Space Complexity: O(logn)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 33.8 MB, less than 5.02%
    public int addDigits(int num) {
        while (num >= 10) {
            num = num / 10 + num % 10;
        }

        return num;
    }

    // Use math formula to judge 9 modulation
    // Time Complexity: O(1)
    // Space Complexity: O(1)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 33.4 MB, less than 5.02%
    public int addDigits2(int num) {
        if (num == 0) {
            return 0;
        }
        if (num % 9 == 0) {
            return 9;
        }
        return num % 9;
    }
}
