/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.1
 */
public class PowerofTwo_231 {
    public static void main(String[] args) {
        int n = -2147483648;
        PowerofTwo_231 mTest = new PowerofTwo_231();
        mTest.isPowerOfTwo(n);
    }

    // Use build-in library
    // Time Complexity: O(1)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 33.7 MB, less than 7.32%
    public boolean isPowerOfTwo(int n) {
        // System.out.println(Integer.bitCount(n));
        return (n > 0) && (Integer.bitCount(n) == 1);
    }

    // Use Iterative
    // Time Complexity: O(1ogn)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 33.5 MB, less than 7.32%
    public boolean isPowerOfTwo2(int n) {
        if (n <= 0) {
            return false;
        }

        while (n % 2 == 0) {
            n /= 2;
        }
        return n == 1;
    }

    // Use Recursive
    // Time Complexity: O(1ogn)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 33.5 MB, less than 7.32%
    public boolean isPowerOfTwo3(int n) {
        return (n > 0) && ((n == 1) || ((n % 2 == 0) && isPowerOfTwo3(n / 2)));
    }

    // Use Bit operation
    // Time Complexity: O(1)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 33.5 MB, less than 7.32%
    public boolean isPowerOfTwo4(int n) {
        return (n > 0) && ((n & (n - 1)) == 0);
    }
}
