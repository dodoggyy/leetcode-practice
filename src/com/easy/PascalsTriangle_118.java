/**
 * 
 */
package com.easy;

import java.util.ArrayList;
import java.util.List;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class PascalsTriangle_118 {
    // Use permutation formula
    // Time Complexity: O(n^2)
    // Space Complexity:O(n^2)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 33.7 MB, less than 5.21%
    public List<List<Integer>> generate(int numRows) {
        List<Integer> mList = new ArrayList<>();
        List<List<Integer>> mResult = new ArrayList<>();
        if (numRows == 0) {
            return mResult;
        }
        for (int i = 0; i < numRows; i++) {
            mList = new ArrayList<>();
            for (int j = 0; j <= i; j++) {
                mList.add(permutation(i, j));
            }
            mResult.add(mList);
        }

        return mResult;
    }

    private int permutation(int mS, int mD) {
        long mAnswer = 1;
        for (int i = 1; i <= mD; i++) {
            mAnswer = mAnswer * (mS - mD + i) / i;
        }

        return (int) mAnswer;
    }

    // Analyze Pascal's Triangle's rule
    // Time Complexity: O(n^2)
    // Space Complexity:O(n^2)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 33.7 MB, less than 5.21%
    public List<List<Integer>> generate2(int numRows) {
        List<List<Integer>> mResult = new ArrayList<>();
        if (numRows == 0) {
            return mResult;
        }

        mResult.add(new ArrayList<>());
        mResult.get(0).add(1);

        for (int mRow = 1; mRow < numRows; mRow++) {
            List<Integer> mList = new ArrayList<>();
            List<Integer> mPreRow = mResult.get(mRow - 1);

            // first row element
            mList.add(1);

            for (int j = 1; j < mRow; j++) {
                mList.add(mPreRow.get(j - 1) + mPreRow.get(j));
            }

            // last row element
            mList.add(1);

            mResult.add(mList);
        }

        return mResult;
    }
}
