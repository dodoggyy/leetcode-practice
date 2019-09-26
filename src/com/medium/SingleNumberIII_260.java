/**
 * 
 */
package com.medium;

import java.util.HashMap;
import java.util.Map;

/**
 * @author Chris Lin
 * @version 1.1
 */
public class SingleNumberIII_260 {

    // Hash Map calculation
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 5 ms, faster than 27.54%
    // Memory Usage: 39.3 MB, less than 41.91%
    public int[] singleNumber(int[] nums) {
        int[] mResult = new int[2];
        int mIndex = 0;

        Map<Integer, Integer> mMap = new HashMap<Integer, Integer>();
        for (int i = 0; i < nums.length; i++) {
            mMap.put(nums[i], mMap.getOrDefault(nums[i], 0) + 1);
        }
        for (int key : mMap.keySet()) {
            if (mMap.get(key) != 2) {
                mResult[mIndex++] = key;

                if (mIndex == 2) {
                    break;
                }
            }
        }
        return mResult;
    }

    // Bitwise operation
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 97.19%
    // Memory Usage: 39.3 MB, less than 36.76%
    public int[] singleNumber2(int[] nums) {
        int[] mResult = new int[2];
        int mDiff = 0;
        for (int num : nums) {
            mDiff ^= num;
        }
        // get least diff bit
        // ex: 6 = 00110b, -6 = 01010b, (6&(-6)) = 10b
        mDiff &= -mDiff;

        for (int num : nums) {
            if ((mDiff & num) == 0) {
                mResult[0] ^= num;
            } else {
                mResult[1] ^= num;
            }
        }

        return mResult;
    }
}
