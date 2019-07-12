/**
 * 
 */
package com.easy;

import java.util.HashSet;
import java.util.Set;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class HappyNumber_202 {
    // 2019年7月11日
    // Use HashSet to keep judgment
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 2 ms, faster than 60.39%
    // Memory Usage: 33.6 MB, less than 6.40%
    public boolean isHappy(int n) {
        if (n < 1) {
            return false;
        }

        Set<Integer> mSet = new HashSet<>();
        while (n != 1 && !mSet.contains(n)) {
            mSet.add(n);
            int mNewN = 0, mTmp = 0;
            while (n > 0) {
                mTmp = n % 10;
                n /= 10;
                mNewN += mTmp * mTmp;
            }
            
            n = mNewN;
        }
        return (n == 1);
    }
}
