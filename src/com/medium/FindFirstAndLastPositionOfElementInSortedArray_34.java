/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class FindFirstAndLastPositionOfElementInSortedArray_34 {
    // 2019年8月10日
    // Use recursive
    // Time Complexity: O(logn)
    // Space Complexity:O(logn)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 43.2 MB, less than 97.87%
    private int[] mResult = { -1, -1 };

    public int[] searchRange(int[] nums, int target) {
        if (nums == null || nums.length == 0) {
            return mResult;
        }
        find(nums, 0, nums.length - 1, target, true);
        find(nums, 0, nums.length - 1, target, false);

        return mResult;
    }

    private void find(int[] nums, int mLeft, int mRight, int target, boolean bIsFindFirst) {
        if (mLeft > mRight) {
            return;
        }
        int mMid = mLeft + (mRight - mLeft) / 2;
        if (nums[mMid] == target) {
            if (bIsFindFirst) {
                if (mResult[0] == -1) {
                    mResult[0] = mMid;
                } else {
                    mResult[0] = Math.min(mResult[0], mMid);
                }
                find(nums, mLeft, mMid - 1, target, bIsFindFirst);
            } else {
                mResult[1] = Math.max(mResult[1], mMid);
                find(nums, mMid + 1, mRight, target, bIsFindFirst);
            }
        } else if (nums[mMid] > target) {
            find(nums, mLeft, mMid - 1, target, bIsFindFirst);
        } else {
            find(nums, mMid + 1, mRight, target, bIsFindFirst);
        }
    }

    // 2019年8月10日
    // Use iterative
    // Time Complexity: O(logn)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 39 MB, less than 100.00%
    public int[] searchRange2(int[] nums, int target) {
        if (nums == null || nums.length == 0) {
            return mResult;
        }
        int mLeft = 0, mRight = nums.length - 1;
        while (mLeft <= mRight) {
            int mMid = (mLeft + mRight) / 2;
            if (nums[mMid] >= target) {
                mRight = mMid - 1;
            } else {
                mLeft = mMid + 1;
            }

            if (nums[mMid] == target) {
                mResult[0] = mMid;
            }
        }

        mLeft = 0;
        mRight = nums.length - 1;
        while (mLeft <= mRight) {
            int mMid = (mLeft + mRight) / 2;
            if (nums[mMid] <= target) {
                mLeft = mMid + 1;
            } else {
                mRight = mMid - 1;
            }

            if (nums[mMid] == target) {
                mResult[1] = mMid;
            }
        }

        return mResult;
    }

}
