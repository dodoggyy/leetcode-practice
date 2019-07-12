/**
 * 
 */
package com.medium;

import java.util.ArrayList;
import java.util.List;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class IntegerBreak_343 {
    // Mathematical induction
    // A(2) = 1 + 1, 1*1=1
    // A(3) = 2 + 1, 2*1=2
    // A(4) = 2 + 2, 2*2=4
    // A(5) = 3 + 2, 3*2=6
    // A(6) = 3 + 3, 3*3=9
    // A(7) = 3 + 4, 3*4=12
    // A(8) = 3 + 3 + 2, 3*3*2 = 18
    // A(9) = 3 + 3 + 3, 3*3*3 = 27
    // A(10)= 3 + 3 + 4, 3*3*4 = 36
    // Time Complexity: O(n/3) = O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 32.8 MB, less than 5.10%
    public int integerBreak(int n) {
        if (n == 2) {
            return 1;
        }
        if (n == 3) {
            return 2;
        }
        int mResult = 1;
        while (n > 4) {
            mResult *= 3;
            n -= 3;
        }
        return n * mResult;
    }

    // Use DP
    // DP[i] = DP[i-3]*3, since i >=7
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 32.9 MB, less than 5.10% 
    public int integerBreak2(int n) {
        List<Integer> mDP = new ArrayList<>();
        mDP.add(0);
        mDP.add(0);
        mDP.add(1);
        mDP.add(2);
        mDP.add(4);
        mDP.add(6);
        mDP.add(9);

        if (n < 7) {
            return mDP.get(n);
        } else {
            for (int i = 7; i <= n; i++) {
                mDP.add(mDP.get(i - 3) * 3);
            }
        }
        return mDP.get(n);
    }
}
