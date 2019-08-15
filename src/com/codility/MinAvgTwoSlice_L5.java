/**
 * 
 */
package com.codility;

/**
 * @author Chris Lin
 * @version 1.0
 * https://app.codility.com/demo/results/trainingV74R85-HMJ/
 */
public class MinAvgTwoSlice_L5 {

    /**
     * @param args
     */
    public static void main(String[] args) {
        // TODO Auto-generated method stub

    }

    // 2019年8月15日
    // Use one pass iterative
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    public int solution(int[] A) {
        // write your code in Java SE 8
        // The trick is
        // that the minimum average slice has "the length of 2 or 3"
        // So only need to calculate the average of the slices of length 2 and 3
        float mResult = Integer.MAX_VALUE;
        int mIndex = 0;
        int mLength = A.length;

        for (int i = 0; i < mLength - 2; i++) {
            float mAvg2 = (float) (A[i] + A[i + 1]) / 2;
            float mAvg3 = (float) (A[i] + A[i + 1] + A[i + 2]) / 3;

            float mMinCurAvg = Math.min(mAvg2, mAvg3);

            if (mMinCurAvg < mResult) {
                mResult = mMinCurAvg;
                mIndex = i;
            }
        }

        float mAvg2 = (A[mLength - 2] + A[mLength - 1]) / 2;
        if (mAvg2 < mResult) {
            mResult = mAvg2;
            mIndex = mLength - 2;
        }
        return mIndex;
    }
}
