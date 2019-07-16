/**
 * 
 */
package com.easy;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class IntersectionOfTwoArraysII_350 {
    // 2019年7月16日
    // Use HashMap to compare
    // Time Complexity: O(n+m)
    // Space Complexity:O(n+m)
    // Runtime: 3 ms, faster than 56.59%
    // Memory Usage: 37.1 MB, less than 90.42%
    public int[] intersect(int[] nums1, int[] nums2) {
        if (nums1 == null || nums2 == null) {
            return null;
        }

        ArrayList<Integer> mList = new ArrayList<>();
        HashMap<Integer, Integer> mMap = new HashMap<>();
        for (int num : nums1) {
            mMap.put(num, mMap.getOrDefault(num, 0) + 1);
        }
        for (int num : nums2) {
            if (mMap.containsKey(num)) {
                int mTmp = mMap.get(num);
                if (mMap.get(num) > 0) {
                    mList.add(num);
                    mMap.put(num, mTmp - 1);
                }
            }
        }

        int[] mResult = new int[mList.size()];
        int mIndex = 0;
        for (int num : mList) {
            mResult[mIndex++] = num;
        }

        return mResult;
    }

    // 2019年7月16日
    // Use merge sort concept (if arrays has been sorted)
    // Time Complexity: O(n+m)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 37.8 MB, less than 50.98%
    public int[] intersect2(int[] nums1, int[] nums2) {
        if (nums1 == null || nums2 == null) {
            return null;
        }
        Arrays.sort(nums1);
        Arrays.sort(nums2);
        int[] mResult = new int[Math.min(nums1.length, nums2.length)];
        int mIndexResult = 0;

        int mIndex1 = 0, mIndex2 = 0;
        while (mIndex1 < nums1.length && mIndex2 < nums2.length) {
            int mNum1 = nums1[mIndex1], mNum2 = nums2[mIndex2];
            if (mNum1 == mNum2) {
                mResult[mIndexResult++] = mNum1;
            }
            if (mNum1 < mNum2) {
                mIndex1++;
            } else if (mNum2 < mNum1) {
                mIndex2++;
            } else { // mNum1 = mNum2
                mIndex1++;
                mIndex2++;
            }
        }

        return Arrays.copyOf(mResult, mIndexResult);
    }
}
