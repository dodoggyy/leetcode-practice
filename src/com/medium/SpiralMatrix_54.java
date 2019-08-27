/**
 * 
 */
package com.medium;

import java.util.ArrayList;
import java.util.List;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class SpiralMatrix_54 {
    // 2019年8月27日
    // Use Simulation
    // Time Complexity: O(m*n)
    // Space Complexity:O(m*n)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 34.7 MB, less than 100.00%
    public List<Integer> spiralOrder(int[][] matrix) {
        List<Integer> mResult = new ArrayList<>();
        if (matrix == null || matrix.length == 0) {
            return mResult;
        }
        int mRow = matrix.length, mColumn = matrix[0].length;
        boolean[][] bIsTraversal = new boolean[mRow][mColumn];

        int[] mDirectionRow = new int[] { 0, 1, 0, -1 };
        int[] mDirectionCol = new int[] { 1, 0, -1, 0 };

        int mCurRow = 0, mCurCol = 0, mDirection = 0;

        for (int i = 0; i < mRow * mColumn; i++) {
            mResult.add(matrix[mCurRow][mCurCol]);
            bIsTraversal[mCurRow][mCurCol] = true;
            int mNextRow = mCurRow + mDirectionRow[mDirection];
            int mNextCol = mCurCol + mDirectionCol[mDirection];

            if (0 <= mNextRow && mNextRow < mRow && 0 <= mNextCol && mNextCol < mColumn
                    && !bIsTraversal[mNextRow][mNextCol]) {
                mCurRow = mNextRow;
                mCurCol = mNextCol;
            } else {
                mDirection = (mDirection + 1) % 4;
                mCurRow += mDirectionRow[mDirection];
                mCurCol += mDirectionCol[mDirection];
            }
        }

        return mResult;
    }

    // 2019年8月27日
    // Use Layer-by-Layer
    // Time Complexity: O(m*n)
    // Space Complexity:O(m*n)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 34.3 MB, less than 100.00%
    public List<Integer> spiralOrder2(int[][] matrix) {
        // [ 1,   1,  1,  1  ]
        // [ 1*,  2,  2,  1' ]
        // [ 1*,  2,  2,  1' ]
        // [ 1*,  1", 1", 1' ]
        List<Integer> mResult = new ArrayList<>();
        if (matrix == null || matrix.length == 0) {
            return mResult;
        }
        int mRow = matrix.length, mColumn = matrix[0].length;

        int mRowStart = 0, mRowEnd = mRow - 1;
        int mColStart = 0, mColEnd = mColumn - 1;
        while (mRowStart <= mRowEnd && mColStart <= mColEnd) {
            for (int i = mColStart; i <= mColEnd; i++) {
                mResult.add(matrix[mRowStart][i]);
            }
            for (int i = mRowStart + 1; i <= mRowEnd; i++) {
                mResult.add(matrix[i][mColEnd]);
            }
            if (mRowStart < mRowEnd && mColStart < mColEnd) {
                for (int i = mColEnd - 1; i > mColStart; i--) {
                    mResult.add(matrix[mRowEnd][i]);
                }
                for (int i = mRowEnd; i > mRowStart; i--) {
                    mResult.add(matrix[i][mColStart]);
                }
            }
            mColStart++;
            mColEnd--;
            mRowStart++;
            mRowEnd--;
        }

        return mResult;
    }
}
