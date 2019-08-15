/**
 * 
 */
package com.codility;

/**
 * @author Chris Lin
 * @version 1.0
 * https://app.codility.com/demo/results/training66ASV2-DGW/
 */
public class TapeEquilibrium_L3 {

    /**
     * @param args
     */
    public static void main(String[] args) {
        // TODO Auto-generated method stub

    }

    // 2019年8月15日
    // Use iterative
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 0.0 MB, less than 100%
    public int solution(int[] A) {
        // write your code in Java SE 8
        int mResult = Integer.MAX_VALUE;
        int mTotal = 0;
        int mLength = A.length;

        for (int i = 0; i < mLength; i++) {
            mTotal += A[i];
        }
        for (int i = 0; i < mLength - 1; i++) {
            mTotal = mTotal - 2 * A[i];
            mResult = Math.min(mResult, Math.abs(mTotal));
        }
        return mResult;
    }
}
