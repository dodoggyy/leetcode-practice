/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class IslandPerimeter_463 {
    // 2019年7月22日
    // Calculate connected case
    // Time Complexity: O(n^2)
    // Space Complexity:O(1)
    // Runtime: 8 ms, faster than 45.04%
    // Memory Usage: 58.7 MB, less than 99.25%
    public int islandPerimeter(int[][] grid) {
        if (grid == null) {
            return 0;
        }
        int mResult = 0;

        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[0].length; j++) {
                if (grid[i][j] == 1) {
                    mResult += helper(grid, i, j);
                }
            }
        }

        return mResult;
    }

    private int helper(int[][] grid, int i, int j) {
        int mCalPerimeter = 0;

        if ((i > 0 && grid[i - 1][j] == 0) || i == 0) {
            mCalPerimeter++;
        }
        if ((j > 0 && grid[i][j - 1] == 0) || j == 0) {
            mCalPerimeter++;
        }
        if ((i + 1 < grid.length && grid[i + 1][j] == 0) || i + 1 == grid.length) {
            mCalPerimeter++;
        }
        if ((j + 1 < grid[0].length && grid[i][j + 1] == 0) || j + 1 == grid[0].length) {
            mCalPerimeter++;
        }

        return mCalPerimeter;
    }

    // 2019年7月22日
    // Use DFS
    // Time Complexity: O(n^2)
    // Space Complexity:O(1)
    // Runtime: 10 ms, faster than 26.66%
    // Memory Usage: 59.2 MB, less than 99.12%
    public int islandPerimeter2(int[][] grid) {
        if (grid == null) {
            return 0;
        }

        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[0].length; j++) {
                if (grid[i][j] == 1) {
                    return DFS(grid, i, j);
                }
            }
        }

        return 0;
    }

    private int DFS(int[][] grid, int i, int j) {
        // judge boundary case
        if (i < 0 || i >= grid.length || j < 0 || j >= grid[0].length) {
            return 1;
        }
        if (grid[i][j] == 0) {
            return 1;
        }
        // if node had been visited will set to -1 and no need to traversal
        if (grid[i][j] == -1) {
            return 0;
        }
        int mCount = 0;

        grid[i][j] = -1;

        mCount += DFS(grid, i - 1, j);
        mCount += DFS(grid, i, j - 1);
        mCount += DFS(grid, i + 1, j);
        mCount += DFS(grid, i, j + 1);

        return mCount;
    }
}
