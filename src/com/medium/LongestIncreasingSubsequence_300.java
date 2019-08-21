/**
 * 
 */
package com.medium;

import java.util.Arrays;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class LongestIncreasingSubsequence_300 {
    // 2019年8月21日
    // Use recursive with memorize
    // Time Complexity: O(n^2)
    // Space Complexity:O(n^2)
    // Runtime: 95 ms, faster than 5.10%
    // Memory Usage: 77.8 MB, less than 5.00%
    public int lengthOfLIS(int[] nums) {
        int mLength = nums.length;
        int[][] mMem = new int[mLength][mLength];
        for (int[] l : mMem) {
            Arrays.fill(l, -1);
        }

        return lengthOfLIS(nums, -1, 0, mMem);
    }

    private int lengthOfLIS(int[] nums, int mIndexPre, int mIndexCur, int[][] mMem) {
        if (mIndexCur == nums.length) {
            return 0;
        }
        if (mMem[mIndexPre + 1][mIndexCur] >= 0) {
            return mMem[mIndexPre + 1][mIndexCur];
        }
        int mTaken = 0;
        if (mIndexPre < 0 || nums[mIndexCur] > nums[mIndexPre]) {
            mTaken = 1 + lengthOfLIS(nums, mIndexCur, mIndexCur + 1, mMem);
        }

        int mNotTaken = lengthOfLIS(nums, mIndexPre, mIndexCur + 1, mMem);
        mMem[mIndexPre + 1][mIndexCur] = Math.max(mTaken, mNotTaken);

        return mMem[mIndexPre + 1][mIndexCur];
    }

    // 2019年8月21日
    // Use DP
    // Time Complexity: O(n^2)
    // Space Complexity:O(n)
    // Runtime: 8 ms, faster than 62.74%
    // Memory Usage: 36.8 MB, less than 57.00%
    public int lengthOfLIS2(int[] nums) {
        // mDP[i] = max(mDP[j]) + 1, for 0 <= j < i
        int mLength = nums.length;
        if (mLength == 0) {
            return 0;
        }
        int[] mDP = new int[mLength];
        mDP[0] = 1;
        int mResult = 1;
        for (int i = 1; i < mLength; i++) {
            int mTmpMax = 0;
            for (int j = 0; j < i; j++) {
                if (nums[i] > nums[j]) {
                    mTmpMax = Math.max(mTmpMax, mDP[j]);
                }
            }
            mDP[i] = mTmpMax + 1;
            mResult = Math.max(mResult, mDP[i]);
        }
        return mResult;
    }

    // 2019年8月21日
    // Use DP with binary search
    // Time Complexity: O(nlogn)
    // Space Complexity:O(n)
    // Runtime: 1 ms, faster than 92.54%
    // Memory Usage: 36.6 MB, less than 100.00%
    public int lengthOfLIS3(int[] nums) {
        int mLength = nums.length;
        if(mLength == 0) {
            return 0;
        }
        int[] mTail = new int[mLength];
        int mResult = 0;
        for(int num: nums) {
            int mIndexLeft = 0, mIndexRight = mResult;
            while(mIndexLeft < mIndexRight) {
                int mIndexMid = (mIndexLeft + mIndexRight)/2;
                if(mTail[mIndexMid] < num) {
                    mIndexLeft = mIndexMid + 1;
                } else {
                    mIndexRight = mIndexMid;
                }
            }
            mTail[mIndexLeft] = num;
            if(mIndexLeft == mResult) {
                mResult++;
            }
        }
        return mResult;
    }
}
