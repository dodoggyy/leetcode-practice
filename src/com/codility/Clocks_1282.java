/**
 * 
 */
package com.codility;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;

/**
 * @author Chris Lin
 * @version 1.0
 *          {@link https://app.codility.com/demo/results/training43Y8VZ-A4K/}
 */
public class Clocks_1282 {

    /**
     * @param args
     */
    public static void main(String[] args) {
        // TODO Auto-generated method stub

    }

    // 2019年8月14日
    // Use HashMap
    // Task Score: 66%
    // Correctness: 50%
    // Performance: 76%
    // Time Complexity: O()
    // Space Complexity:O()
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 0.0 MB, less than 100%
    public int solution(int[][] A, int P) {
        // write your code in Java SE 8
        int mResult = 0;
        int N = A.length;
        int M = A[0].length;
        HashMap<List<Integer>, Integer> mMap = new HashMap<>();
        for (int i = 0; i < N; i++) {
            int[] mTmpSlice = new int[M + 1];
            calInterval(A, i, M, P, mTmpSlice);
            List<Integer> mList = new ArrayList<>();
            for (int slice : mTmpSlice) {
                mList.add(slice);
            }
            mMap.put(mList, mMap.getOrDefault(mList, 0) + 1);
        }
        for (List<Integer> key : mMap.keySet()) {
            int mValue = mMap.get(key);
            if (mValue >= 2) {
                mResult += mValue * (mValue - 1) / 2;
            }
        }

        return mResult;
    }

    private void calInterval(int[][] A, int mIndex, int M, int P, int[] mTmpSlice) {
        Arrays.sort(A[mIndex]);
        int mTmpTotal = 0;
        for (int i = 0; i < M - 1; i++) {
            mTmpSlice[i] = A[mIndex][i + 1] - A[mIndex][i];
            mTmpTotal += mTmpSlice[i];
        }

        mTmpSlice[M] = P - mTmpTotal;

        Arrays.sort(mTmpSlice);
    }

}
