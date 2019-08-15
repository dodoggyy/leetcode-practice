/**
 * 
 */
package com.codility;

/**
 * @author Chris Lin
 * @version 1.0 https://app.codility.com/demo/results/trainingXDPAXX-WC4/
 * 
 */
public class MinMaxDivision_L14 {

    /**
     * @param args
     */
    public static void main(String[] args) {
        // TODO Auto-generated method stub

    }

    // 2019年8月16日
    // Use binary search
    // Time Complexity: O(N*log(N+M))
    // Space Complexity:O(1)
    public int solution(int K, int M, int[] A) {
        // write your code in Java SE 8
        int mMin = 0;
        int mMax = 0;
        for (int num : A) {
            mMax += num;
            mMin = Math.max(mMin, num);
        }
        int mResult = mMax;
        // binary search to find minMaxSum
        while (mMin <= mMax) {
            int mMid = (mMin + mMax) / 2;
            if (helper(A, mMid, K)) {
                mMax = mMid - 1;
                mResult = mMid;
            } else { // not OK and try bigger
                mMin = mMid + 1;
            }
        }
        return mResult;
    }

    private boolean helper(int[] A, int mMid, int K) {
        int mSum = 0;
        for (int num : A) {
            mSum += num;
            if (mSum > mMid) {
                mSum = num; // next block
                K--;
            }
            if (K == 0) { // cannot achieve minMaxSum = mMid
                return false;
            }
        }
        return true;
    }
}
