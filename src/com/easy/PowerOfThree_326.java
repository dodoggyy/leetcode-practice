/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.1
 */
public class PowerOfThree_326 {
    // 2019年7月15日
    // Judge divided by 3 can be modulus
    // Time Complexity: O(logn)
    // Space Complexity:O(1)
    // Runtime: 11 ms, faster than 67.66%
    // Memory Usage: 35.8 MB, less than 5.55%
    public boolean isPowerOfThree(int n) {
        if (n <= 0) {
            return false;
        }

        while (n != 1) {
            if (n % 3 != 0) {
                return false;
            }
            n /= 3;
        }
        return true;
    }

    // 2019年7月15日
    // Use Math formula
    // log3(n) = logb(n)/logb(3)
    // Time Complexity: O(logn)
    // Space Complexity:O(1)
    // Runtime: 12 ms, faster than 16.79%
    // Memory Usage: 36 MB, less than 5.01%
    public boolean isPowerOfThree2(int n) {
        return Math.log10(n) / Math.log10(3) % 1 == 0;
    }

    // 2019年9月26日
    // Use Recursive
    // Time Complexity: O(logn)
    // Space Complexity:O(1)
    // Runtime: 11 ms, faster than 78.87%
    // Memory Usage: 35.4 MB, less than 9.38%
    public boolean isPowerOfThree3(int n) {
        return (n > 0) && ((n == 1) || (n % 3 == 0 && isPowerOfThree3(n / 3)));
    }

    // 2019年9月26日
    // Simplify Judge divided by 3 can be modulus
    // Time Complexity: O(logn)
    // Space Complexity:O(1)
    // Runtime: 11 ms, faster than 78.87%
    // Memory Usage: 35.9 MB, less than 6.25%
    public boolean isPowerOfThree4(int n) {
        if (n <= 0) {
            return false;
        }
        while (n % 3 == 0) {
            n /= 3;
        }

        return (n == 1);
    }
}
