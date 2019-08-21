/**
 * 
 */
package com.medium;

import java.util.Arrays;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class PerfectSquares_279 {
    // 2019年8月21日
    // Use Lagrange's four-square theorem
    // Time Complexity: O(n^(1/2))
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 33.5 MB, less than 94.44%
    public int numSquares(int n) {
        while (n % 4 == 0) {
            n /= 4;
        }
        if (n % 8 == 7) {
            return 4;
        }
        for (int num1 = 0; num1 * num1 < n; num1++) {
            int num2 = (int) Math.sqrt(n - num1 * num1);
            if (num1 * num1 + num2 * num2 == n) {
                return (num1 == 0 || num2 == 0) ? 1 : 2;
            }
        }
        return 3;
    }

    // 2019年8月21日
    // Use DP
    // Time Complexity: O(n^2)
    // Space Complexity:O(n)
    // Runtime: 20 ms, faster than 75.37%
    // Memory Usage: 35.2 MB, less than 18.05%
    public int numSquares2(int n) {
        if (n < 0) {
            return 0;
        }

        int[] mCountPerfectSqrt = new int[n + 1];
        Arrays.fill(mCountPerfectSqrt, Integer.MAX_VALUE);
        mCountPerfectSqrt[0] = 0;
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j * j <= i; j++) {
                mCountPerfectSqrt[i] = Math.min(mCountPerfectSqrt[i], mCountPerfectSqrt[i - j * j] + 1);
            }
        }

        return mCountPerfectSqrt[n];
    }
}
