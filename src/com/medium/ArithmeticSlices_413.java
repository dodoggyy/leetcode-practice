/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ArithmeticSlices_413 {
    // Use DP
    // Reduce the problem to all 1 sub arrays
    // mDP[i-2] = bIsSlice(A[i], A[i+1], A[i+2])
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 35.5 MB, less than 100.00%
    public int numberOfArithmeticSlices(int[] A) {
        if (A.length < 3) {
            return 0;
        }
        int[] mDP = new int[A.length - 2];
        for (int i = 2; i < A.length; i++) {
            if ((A[i - 1] - A[i - 2]) == A[i] - A[i - 1]) {
                mDP[i - 2] = 1;
            }
        }
        return helper(mDP);
    }

    // number of arrays values
    private int helper(int[] mDP) {
        int mSum = 0, mCurrent = 0;
        for (int i = 0; i < mDP.length; i++) {
            if (mDP[i] > 0) {
                mSum += ++mCurrent;
            } else {
                mCurrent = 0;
            }
        }
        return mSum;
    }

    // Use DP
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 35.4 MB, less than 100.00%
    public int numberOfArithmeticSlices2(int[] A) {
        int mSum = 0, mCurrent = 0;
        for (int i = 2; i < A.length; i++) {
            if ((A[i - 1] - A[i - 2]) == A[i] - A[i - 1]) {
                mSum += ++mCurrent;
            } else {
                mCurrent = 0;
            }
        }
        return mSum;
    }
}
