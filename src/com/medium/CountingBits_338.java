/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class CountingBits_338 {

    /**
     * @param args
     */
    public static void main(String[] args) {
        // TODO Auto-generated method stub
    }

    // simple method
    // Runtime: 1 ms, faster than 99.86%
    // Memory Usage: 36.9 MB, less than 100.00%
    public int[] countBits(int num) {
        int[] mResult = new int[num + 1];
        for (int i = 0; i <= num; i++) {
            mResult[i] = Integer.bitCount(i);
        }
        return mResult;
    }

    // use DP
    // Runtime: 1 ms, faster than 99.86%
    // Memory Usage: 36.5 MB, less than 100.00%
    public int[] countBits2(int num) {
        int[] mResult = new int[num + 1];
        for (int i = 1; i <= num; i++) {
            //mResult[i] = mResult[i & (i - 1)] + 1; // judge 2^n
            mResult[i] = mResult[(i >> 1)] + (i&1); //e.g. DP[15] = DP[7] + (15&1) 
        }
        return mResult;
    }
}
