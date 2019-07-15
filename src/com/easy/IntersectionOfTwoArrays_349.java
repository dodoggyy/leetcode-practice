/**
 * 
 */
package com.easy;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashSet;
import java.util.Set;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class IntersectionOfTwoArrays_349 {
    // 2019年7月15日
    // Use HashSet to compare
    // Time Complexity: O(n + m)
    // Space Complexity:O(n + m)
    // Runtime: 2 ms, faster than 98.27%
    // Memory Usage: 36.7 MB, less than 89.66%
    public int[] intersection(int[] nums1, int[] nums2) {
        Set<Integer> mSet = new HashSet<>();
        ArrayList<Integer> mList = new ArrayList<>();
        for (int num : nums1) {
            mSet.add(num);
        }
        for (int num : nums2) {
            if (mSet.contains(num)) {
                mList.add(num);
                mSet.remove(num);
            }
        }
        int[] mResult = new int[mList.size()];
        int mIndex = 0;
        for (int num : mList) {
            mResult[mIndex++] = num;
        }

        return mResult;

        // Use below function may cause lots of time in LeetCode
        // return mList.stream().mapToInt(Integer::valueOf).toArray();
    }

    // 2019年7月15日
    // Use merge sort concept (if arrays has been sorted)
    // Time Complexity: O(n + m)
    // Space Complexity:O(1)
    // Runtime: 2 ms, faster than 98.27%
    // Memory Usage: 36.6 MB, less than 89.66%
    public int[] intersection2(int[] nums1, int[] nums2) {
        ArrayList<Integer> mList = new ArrayList<>();
        Arrays.sort(nums1);
        Arrays.sort(nums2);

        int mIndex1 = 0, mIndex2 = 0;
        while (mIndex1 < nums1.length && mIndex2 < nums2.length) {
            int mNum1 = nums1[mIndex1], mNum2 = nums2[mIndex2];
            if (mNum1 == mNum2) {
                mList.add(mNum1);
            }
            if (mNum1 < mNum2) {
                while (mIndex1 < nums1.length && mNum1 == nums1[mIndex1]) {
                    mIndex1++;
                }
            } else {
                while (mIndex2 < nums2.length && mNum2 == nums2[mIndex2]) {
                    mIndex2++;
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
}
