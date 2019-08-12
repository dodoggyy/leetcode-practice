/**
 * 
 */
package com.medium;

import java.util.HashMap;
import java.util.Map;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class SingleNumberII_137 {
    // 2019年8月12日
    // Use hash map
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 6 ms, faster than 11.57%
    // Memory Usage: 37.3 MB, less than 100.00%
    public int singleNumber(int[] nums) {
        Map<Integer, Integer> mMap = new HashMap<>();
        for (int num : nums) {
            mMap.put(num, mMap.getOrDefault(num, 0) + 1);
        }
        for (int num : mMap.keySet()) {
            if (mMap.get(num) == 1) {
                return num;
            }
        }
        return -1;
    }

    // 2019年8月12日
    // Use bit operation
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 37.3 MB, less than 100.00%
    public int singleNumber2(int[] nums) {
        int mSeenOnce = 0, mSeenTwice = 0;

        for (int num : nums) {
            mSeenOnce = ~mSeenTwice & (mSeenOnce ^ num);
            mSeenTwice = ~mSeenOnce & (mSeenTwice ^ num);
        }

        return mSeenOnce;
    }

    // 2019年8月12日
    // Use bit operation 2
    // Time Complexity: O(32*n) = O(n)
    // Space Complexity:O(1)
    // Runtime: 2 ms, faster than 60.93%
    // Memory Usage: 36.9 MB, less than 100.00%
    public int singleNumber3(int[] nums) {
        int mResult = 0;

        for (int i = 0; i < 32; i++) {
            int mCount = 0;
            int mBit = 1 << i;
            for (int num : nums) {
                if ((num & mBit) != 0) {
                    mCount++;
                }
            }
            if (mCount % 3 != 0) {
                mResult |= mBit;
            }
        }
        return mResult;
    }
}
