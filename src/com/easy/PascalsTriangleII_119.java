/**
 * 
 */
package com.easy;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class PascalsTriangleII_119 {
    // 2019年7月10日
    // Use permutation formula
    // Time Complexity: O(n^2)
    // Space Complexity:O(n)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 33.6 MB, less than 5.08%
    public List<Integer> getRow(int rowIndex) {
        List<Integer> mResult = new ArrayList<>();

        for (int i = 0; i <= rowIndex; i++) {
            mResult.add(permutation(rowIndex, i));
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

    // 2019年7月10日
    // Analyze Pascal's Triangle's rule
    // Time Complexity: O(n^2)
    // Space Complexity:O(n)
    // Runtime: 1 ms, faster than 88.92%
    // Memory Usage: 33.9 MB, less than 5.08%
    public List<Integer> getRow2(int rowIndex) {
        Integer[] mResult = new Integer[rowIndex + 1];
        Arrays.fill(mResult, 0);
        mResult[0] = 1;

        for (int i = 1; i <= rowIndex; i++) {
            for (int j = i; j > 0; j--) {
                mResult[j] += mResult[j - 1];
            }
        }

        return Arrays.asList(mResult);
    }
}
