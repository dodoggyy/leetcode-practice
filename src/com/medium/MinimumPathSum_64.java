/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class MinimumPathSum_64 {
    // Recursion with memorize
    // minSum(x,y) = grid[y][x] + min(minSum(x-1, y), ,minSum(x, y-1))
    // Time Complexity: O(m*n)
    // Space Complexity:O(m*n)
    // Runtime: 1 ms, faster than 99.61%
    // Memory Usage: 39.2 MB, less than 91.10%
    private int[][] mMinSum;

    public int minPathSum(int[][] grid) {
        int m = grid.length;
        if (m == 0) {
            return 0;
        }
        int n = grid[0].length;
        mMinSum = new int[m][n];

        return minPathSum(grid, m - 1, n - 1, m, n);
    }

    private int minPathSum(int[][] grid, int mIndexRow, int mIndexColumn, int n, int m) {
        if (mIndexRow == 0 && mIndexColumn == 0) {
            return grid[mIndexRow][mIndexColumn];
        }
        if (mIndexRow < 0 || mIndexColumn < 0) {
            return Integer.MAX_VALUE;
        }
        if (mMinSum[mIndexRow][mIndexColumn] > 0) {
            return mMinSum[mIndexRow][mIndexColumn];
        }

        mMinSum[mIndexRow][mIndexColumn] = grid[mIndexRow][mIndexColumn]
                + Math.min(minPathSum(grid, mIndexRow - 1, mIndexColumn, n, m),
                        minPathSum(grid, mIndexRow, mIndexColumn - 1, n, m));

        return mMinSum[mIndexRow][mIndexColumn];
    }

    // Use DP
    // dp(x,y) = grid[y][x] + min(dp(x-1, y), ,dp(x, y-1))
    // Time Complexity: O(m*n)
    // Space Complexity: O(1)
    // Runtime: 2 ms, faster than 90.23%
    // Memory Usage: 42.7 MB, less than 59.97%
    public int minPathSum2(int[][] grid) {
        int m = grid.length;
        if (m == 0) {
            return 0;
        }
        int n = grid[0].length;

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (i == 0 && j == 0) {
                    continue; // do nothing
                } else if (i == 0) { // [0][j] case
                    grid[i][j] += grid[i][j - 1];
                } else if (j == 0) { // [i][0] case
                    grid[i][j] += grid[i - 1][j];
                } else { // [i][j] case
                    grid[i][j] += Math.min(grid[i - 1][j], grid[i][j - 1]);
                }
            }
        }

        return grid[m - 1][n - 1];
    }
}
