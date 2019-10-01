/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class Sqrtx_69 {
    // 2019年10月1日
    // Use binary search
    // Time Complexity: O(logn)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 33.8 MB, less than 5.00%
    public int mySqrt(int x) {
        if (x <= 1) {
            return x;
        }
        int mIndexLow = 1, mIndexHigh = x;
        while (mIndexLow <= mIndexHigh) {
            int mIndexMid = mIndexLow + (mIndexHigh - mIndexLow) / 2;
            if (mIndexMid > x / mIndexMid) {
                mIndexHigh = mIndexMid - 1;
            } else if (mIndexMid <= x / mIndexMid) {
                mIndexLow = mIndexMid + 1;
            }
        }
        return mIndexHigh;
    }

    // 2019年10月1日
    // Use brute force
    // Time Complexity: O(n^1/2)
    // Space Complexity:O(1)
    // Runtime: 16 ms, faster than 8.48%
    // Memory Usage: 33.6 MB, less than 5.00%
    public int mySqrt2(int x) {
        if (x <= 1) {
            return x;
        }
        for (long i = 1; i <= x; i++) {
            long mSqure = i * i;
            if (mSqure > x) {
                return (int) (i - 1);
            }
        }
        
        // divide method
        /*
        for (int i = 1; i <= x; i++) {
          if (i > x/i) {
              return (int) (i - 1);
          }
        }
        */
        return -1;
    }

    // 2019年10月1日
    // Use Newton Method
    // Time Complexity: O(logn)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 33.7 MB, less than 5.00%
    public int mySqrt3(int x) {
        // f(x) = x^2 - a = 0, x = sqrt(a)
        // f'(x) = 2*x
        // x0 = a
        // Xn+1 = Xn - (Xn^2 - a) / 2*Xn
        //      = (Xn + a / Xn) / 2
        
        long mResult = x;
        while(mResult * mResult > x) {
            mResult = (mResult + x / mResult) / 2;
        }
        
        return (int)mResult;
    }
}
