/**
 * 
 */
package com.easy;

/**
 * @author linquanliang
 * @version 1.0
 */
public class CountNegativeNumbersInASortedMatrix_1351 {
    // 2020年2月20日
    // Use linear search
    // Time Complexity: O(m*n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 41.3 MB, less than 100.00%
    public int countNegatives(int[][] grid) {
        int NonNegCount = 0;
        int m = grid.length, n = grid[0].length;

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] < 0) {
                    break;
                }
                NonNegCount++;
            }
        }

        return m * n - NonNegCount;
    }

    // 2020年2月20日
    // Use binary search
    // Time Complexity: O(mlogn)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 41.8 MB, less than 100.00%
    public int countNegatives2(int[][] grid) {
        int negCount = 0;
        int m = grid.length, n = grid[0].length;
        for (int i = 0; i < m; i++) {
            int index = negPosStart(grid[i]);
            if (index < 0) {
                continue;
            }
            negCount += n - index;
            // [1 1 1 -1 -1] , 5 ,3
        }

        return negCount;
    }

    public int negPosStart(int arr[]) {
        int start = 0, end = arr.length - 1;

        while (start <= end) {
            int mid = start + (end - start) / 2;
            if (arr[mid] < 0) {
                if (mid == 0 || arr[mid - 1] >= 0) {
                    return mid;
                } else {
                    end = mid - 1;
                }
            } else {
                start = mid + 1;
            }
        }

        return -1;
    }
}
