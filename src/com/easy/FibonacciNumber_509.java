/**
 * 
 */
package com.easy;

import java.util.HashMap;
import java.util.Map;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class FibonacciNumber_509 {
    // 2019年7月25日
    // Use Iterative
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 32.9 MB, less than 5.14%
    public int fib(int N) {
        if (N == 0) {
            return 0;
        }
        if (N == 1) {
            return 1;
        }
        int mFibNMinus2 = 0, mFibNMinus1 = 1;
        int mFibN = 0;
        for (int i = 2; i <= N; i++) {
            mFibN = mFibNMinus2 + mFibNMinus1;
            mFibNMinus2 = mFibNMinus1;
            mFibNMinus1 = mFibN;
        }
        return mFibN;
    }

    // 2019年7月25日
    // Use Recursive without memorize
    // Time Complexity: O(2^n)
    // Space Complexity:O(n)
    // Runtime: 9 ms, faster than 25.46%
    // Memory Usage: 33 MB, less than 5.14%
    public int fib2(int N) {
        if (N == 0) {
            return 0;
        }
        if (N == 1) {
            return 1;
        }
        return fib2(N - 1) + fib2(N - 2);
    }

    // 2019年7月25日
    // Use Recursive with memorize
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 1 ms, faster than 37.02%
    // Memory Usage: 32.9 MB, less than 5.14%
    private Map<Integer, Integer> mMap = new HashMap<>();

    public int fib3(int N) {
        if (N == 0) {
            return 0;
        }
        if (N == 1) {
            return 1;
        }
        mMap.put(0, 0);
        mMap.put(1, 1);
        if (mMap.containsKey(N)) {
            return mMap.get(N);
        } else {
            int mResult = fib3(N - 1) + fib3(N - 2);
            mMap.put(N, mResult);
            return mResult;
        }
    }
}
