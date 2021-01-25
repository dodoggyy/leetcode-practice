/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.1
 */
public class ClimbingStairs_70 {

    // Use Fibonacci number concept
    // Time Complexity: O(n)
    // Space complexity: O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 31.8 MB, less than 100.00%
    public int climbStairs(int n) {
        if (n <= 2) {
            return n;
        }
        int fibNminus2 = 1, fibNminus1 = 2;
        for (int i = 2; i < n; i++) {
            int fibN = fibNminus2 + fibNminus1;
            fibNminus2 = fibNminus1;
            fibNminus1 = fibN;
        }

        return fibNminus1;
    }

    // Use Recursive with memorize
    // Time Complexity: O(n)
    // Space complexity: O(n)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 33 MB, less than 5.26%
    public int climbStairs2(int n) {
        int[] mDP = new int[n + 1];

        return helper(0, n, mDP);
    }

    private int helper(int i, int n, int mDP[]) {
        if (i > n) {
            return 0;
        }
        if (i == n) {
            return 1;
        }
        if (mDP[i] != 0) {
            return mDP[i];
        }
        mDP[i] = helper(i + 1, n, mDP) + helper(i + 2, n, mDP);
        return mDP[i];
    }

    // Use DP
    // Time Complexity: O(n)
    // Space complexity: O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 33 MB, less than 5.26%
    public int climbStairs3(int n) {
        if(n == 1) {
            return 1;
        }
        int[] mDP = new int[n+1];
        mDP[1] = 1;
        mDP[2] = 2;
        for(int i = 3; i <= n; i++) {
            mDP[i] = mDP[i - 1] + mDP[i-2];
        }
        return mDP[n];
    }
}
