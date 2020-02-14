/**
 * 
 */
package com.easy;

import java.util.HashSet;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class SetMismatch_645 {
    // 2020年2月13日
    // Use swap index:
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 2 ms, faster than 95.10%
    // Memory Usage: 42.8 MB, less than 14.29%
    public int[] findErrorNums(int[] nums) {
        for (int i = 0; i < nums.length; i++) {
            while (nums[i] != (i + 1) && nums[nums[i] - 1] != nums[i]) {
                swap(nums, i, nums[i] - 1);
            }
        }
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] != (i + 1)) {
                return new int[] { nums[i], i + 1 };
            }
        }

        return null;
    }

    public void swap(int[] nums, int i, int j) {
        int tmp = nums[i];
        nums[i] = nums[j];
        nums[j] = tmp;
    }

    // 2020年2月13日
    // Use Hash Set:
    // Time Complexity: O(2n) = O(n)
    // Space Complexity:O(n)
    // Runtime: 14 ms, faster than 33.39%
    // Memory Usage: 42.9 MB, less than 14.29%
    public int[] findErrorNums2(int[] nums) {
        int[] mResult = new int[2];
        HashSet<Integer> mSet = new HashSet<>();

        for (int num : nums) {
            if (mSet.contains(num)) {
                mResult[0] = num;
            }
            mSet.add(num);
        }
        for (int i = 0; i < nums.length; i++) {
            if (mSet.contains(i + 1)) {
                mSet.remove(i + 1);
            } else {
                mResult[1] = i + 1;
                break;
            }
        }

        return mResult;
    }
}
