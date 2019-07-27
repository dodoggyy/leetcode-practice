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
public class KDiffPairsInAnArray_532 {
    // 2019年7月27日
    // Use HashMap record
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 9 ms, faster than 80.95%
    // Memory Usage: 39.8 MB, less than 86.71%
    public int findPairs(int[] nums, int k) {
        if (nums == null || nums.length == 0 || k < 0) {
            return 0;
        }
        int mResult = 0;
        Map<Integer, Integer> mMap = new HashMap<>();
        for (int i : nums) {
            if (mMap.containsKey(i)) {
                if (k == 0 && mMap.get(i) == 1) {
                    mResult++;
                }
                mMap.put(i, mMap.get(i) + 1);
            } else {
                if (mMap.containsKey(i - k)) {
                    mResult++;
                }
                if (mMap.containsKey(i + k)) {
                    mResult++;
                }
                mMap.put(i, 1);
            }
        }

        return mResult;
    }
}
