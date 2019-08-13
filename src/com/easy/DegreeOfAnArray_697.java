/**
 * 
 */
package com.easy;

import java.util.Collections;
import java.util.HashMap;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class DegreeOfAnArray_697 {
    // 2019年8月14日
    // Use hash map
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 29 ms, faster than 35.00%
    // Memory Usage: 41.1 MB, less than 88.89%
    public int findShortestSubArray(int[] nums) {
        HashMap<Integer, Integer> mMapCount = new HashMap<>();
        HashMap<Integer, Integer> mMapLeft = new HashMap<>();
        HashMap<Integer, Integer> mMapRight = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            mMapCount.put(nums[i], mMapCount.getOrDefault(nums[i], 0) + 1);
            if (!mMapLeft.containsKey(nums[i])) {
                mMapLeft.put(nums[i], i);
            }
            mMapRight.put(nums[i], i);
        }
        int mResult = nums.length;
        int mDegree = Collections.max(mMapCount.values());
        System.out.println(mDegree);
        for (int key : mMapCount.keySet()) {
            if (mMapCount.get(key) == mDegree) {
                mResult = Math.min(mResult, mMapRight.get(key) - mMapLeft.get(key) + 1);
            }
        }

        return mResult;
    }
}
