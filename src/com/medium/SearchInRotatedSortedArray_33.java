/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class SearchInRotatedSortedArray_33 {
    // 2019年8月10日
    // Use iterative
    // Time Complexity: O(logn)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 39.5 MB, less than 16.98%
    public int search(int[] nums, int target) {
        int mLeft = 0, mRight = nums.length - 1;
        while (mLeft <= mRight) {
            int mMid = (mLeft + mRight) / 2;
            if (nums[mMid] == target) {
                return mMid;
            }

            if (nums[mMid] < nums[mRight]) { // check left or right part has
                                             // order
                if (nums[mMid] < target && nums[mRight] >= target) {
                    mLeft = mMid + 1;
                } else {
                    mRight = mMid - 1;
                }
            } else {
                if (nums[mMid] > target && nums[mLeft] <= target) {
                    mRight = mMid - 1;
                } else {
                    mLeft = mMid + 1;
                }
            }
        }

        return -1;
    }

    // 2019年8月10日
    // Use recursive
    // Time Complexity: O(logn)
    // Space Complexity:O(1ogn)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 40.7 MB, less than 5.66%
    public int search2(int[] nums, int target) {
        return search(nums, 0, nums.length - 1, target);
    }

    private int search(int[] nums, int mLeft, int mRight, int target) {
        int mMid = (mLeft + mRight) / 2;
        if (mLeft > mRight) {
            return -1;
        }
        if (nums[mMid] == target) {
            return mMid;
        }
        if (nums[mMid] < nums[mRight]) {
            if (target > nums[mMid] && target <= nums[mRight]) {
                return search(nums, mMid + 1, mRight, target);
            } else {
                return search(nums, mLeft, mMid - 1, target);
            }
        } else {
            if (target < nums[mMid] && target >= nums[mLeft]) {
                return search(nums, mLeft, mMid - 1, target);
            } else {
                return search(nums, mMid + 1, mRight, target);
            }
        }
    }
}
