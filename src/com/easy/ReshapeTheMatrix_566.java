/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ReshapeTheMatrix_566 {
    // 2019年8月1日
    // Use iterative
    // Time Complexity: O(mn)
    // Space Complexity:O(mn)
    // m:row length, n:column length
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 38.8 MB, less than 100.00%
    public int[][] matrixReshape(int[][] nums, int r, int c) {
        if (nums == null) {
            return null;
        }
        if (r * c != nums.length * nums[0].length) {
            return nums;
        }
        int[][] mResult = new int[r][c];
        int mIndexRow = 0, mIndexColum = 0;
        for (int i = 0; i < r; i++) {
            for (int j = 0; j < c; j++) {
                mResult[i][j] = nums[mIndexRow][mIndexColum];
                if ((mIndexColum + 1) == nums[0].length) {
                    mIndexColum = 0;
                    mIndexRow++;
                } else {
                    mIndexColum++;
                }
            }
        }

        return mResult;
    }

    // 2019年8月1日
    // Use iterative with optimize code
    // Time Complexity: O(mn)
    // Space Complexity:O(mn)
    // m:row length, n:column length
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 38.8 MB, less than 100.00%
    public int[][] matrixReshape2(int[][] nums, int r, int c) {
        int mRowLength = nums.length, mColumnLength = nums[0].length;
        if (r * c != mRowLength * mColumnLength) {
            return nums;
        }
        int[][] mResult = new int[r][c];
        for (int i = 0; i < r * c; i++) {
            mResult[i / c][i % c] = nums[i / mColumnLength][i % mColumnLength];
        }
        return mResult;
    }
}
