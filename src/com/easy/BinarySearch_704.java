/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class BinarySearch_704 {

    // Iterative version
    // Runtime: 2 ms, faster than 99.94%
    // Memory Usage: 40.8 MB, less than 39.04%
    public int search(int[] nums, int target) {
        int mIndexLeft = 0;
        int mIndexRight = nums.length - 1;

        while (mIndexLeft <= mIndexRight) {
            int mIndexMid = (mIndexLeft + mIndexRight) / 2;
            if (target == nums[mIndexMid]) {
                return mIndexMid;
            } else if (target > nums[mIndexMid]) {
                mIndexLeft = mIndexMid + 1;
            } else if (target < nums[mIndexMid]) {
                mIndexRight = mIndexMid - 1;
            }
        }

        return -1;
    }

    // Recursive version
    // Runtime: 2 ms, faster than 99.94%
    // Memory Usage: 40.9 MB, less than 10.33%
    public int search2(int[] nums, int target) {
        int mIndexLeft = 0;
        int mIndexRight = nums.length - 1;

        return searchRecursive(nums, target, mIndexLeft, mIndexRight);
    }

    public static int searchRecursive(int[] nums, int target, int mIndexLeft, int mIndexRight) {
        int mIndexMid = (mIndexLeft + mIndexRight) / 2;

        if (mIndexLeft > mIndexRight) {
            return -1;
        }

        if (target == nums[mIndexMid]) {
            return mIndexMid;
        } else if (target > nums[mIndexMid]) {
            return searchRecursive(nums, target, (mIndexMid + 1), mIndexRight);
        } else { // (target < nums[mIndexMid])
            return searchRecursive(nums, target, mIndexLeft, (mIndexMid - 1));
        }
    }
}
