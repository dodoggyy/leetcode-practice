/**
 * 
 */
package com.easy;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Stack;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class NextGreaterElementI_496 {
    // 2019年7月23日
    // Use Stack & HashMap
    // Time Complexity: O()
    // Space Complexity:O(m)
    // m: length of nums1. n: length of nums2
    // Runtime: 3 ms, faster than 79.16%
    // Memory Usage: 38 MB, less than 95.54%
    public int[] nextGreaterElement(int[] nums1, int[] nums2) {
        if (nums1 == null || nums1.length == 0 || nums2 == null || nums2.length == 0) {
            return new int[] {};
        }

        int mLength1 = nums1.length, mLength2 = nums2.length;
        int[] mResult = new int[mLength1];

        HashMap<Integer, Integer> mMap = new HashMap<>();
        Stack<Integer> mStack = new Stack<>();

        for (int i = 0; i < mLength2; i++) {
            while (!mStack.isEmpty() && mStack.peek() < nums2[i]) {
                mMap.put(mStack.pop(), nums2[i]);
            }
            mStack.push(nums2[i]);
        }
        for (int i = 0; i < mLength1; i++) {
            mResult[i] = mMap.getOrDefault(nums1[i], -1);
        }

        return mResult;
    }

    // 2019年7月23日
    // Use brute force
    // Time Complexity: O(m*n)
    // Space Complexity:O(m)
    // m: length of nums1. n: length of nums2
    // Runtime: 10 ms, faster than 7.84%
    // Memory Usage: 37.2 MB, less than 98.85%
    public int[] nextGreaterElement2(int[] nums1, int[] nums2) {
        int[] mResult = new int[nums1.length];
        Arrays.fill(mResult, -1);

        for (int i = 0; i < nums1.length; i++) {
            int j = 0, k = 0;
            for (j = 0; j < nums2.length; j++) {
                if (nums1[i] == nums2[j]) {
                    break;
                }
            }
            for (k = j + 1; k < nums2.length; k++) {
                if (nums2[k] > nums2[j]) {
                    mResult[i] = nums2[k];
                    break;
                }
            }

        }

        return mResult;
    }
}
