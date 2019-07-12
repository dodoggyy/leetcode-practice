/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class UniquePaths_62 {
    // Use DP
    // uniquePaths(m,n) = dp[n - 1][m] + dp[n][m-1]
    // Time Complexity: O(n*m)
    // Space Complexity:O(n*m)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 32.9 MB, less than 5.00%
    public int uniquePaths(int m, int n) {
        if (m == 0 || n == 0) {
            return 1;
        }
        int[][] mDP = new int[n][m];
        mDP[0][0] = 1;
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                if (i == 0 && j == 0) {
                    continue;
                }
                if (i == 0) {
                    mDP[i][j] = mDP[i][j - 1];
                } else if (j == 0) {
                    mDP[i][j] = mDP[i - 1][j];
                } else {
                    mDP[i][j] = mDP[i][j - 1] + mDP[i - 1][j];
                }
            }
        }
        return mDP[n - 1][m - 1];
    }

    // Use Permutation formula
    // uniquePaths(m,n) = C((m+n-2), (m-1))
    // Time Complexity: O(1)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 33 MB, less than 5.00%
    public int uniquePaths2(int m, int n) {
        long mAnswer = 1;
        // C(S, D)
        int S = m + n - 2;
        int D = n - 1;
        for (int i = 1; i <= D; i++) {
            mAnswer = mAnswer * (S - D + i) / i;
        }

        return (int) mAnswer;
    }
}
