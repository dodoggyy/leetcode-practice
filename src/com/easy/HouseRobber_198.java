/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.1
 */
public class HouseRobber_198 {

    // DP[i] = max(DP[i-2] + nums[i], DP[i-1])
    // Time Complexity: O(n)
    // Space complexity: O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 33.5 MB, less than 100.00%
    public int rob(int[] nums) {
        int mRob = 0;
        int mNotRob = 0;
        for (int i = 0; i < nums.length; i++) {
            int mCurrent = Math.max(mNotRob + nums[i], mRob);
            mNotRob = mRob;
            mRob = mCurrent;
        }

        return mRob;
    }

    // Judge even & odd position
    // Time Complexity: O(n)
    // Space complexity: O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 33.2 MB, less than 100.00%
    public int rob2(int[] nums) {
        int mEvenRob = 0, mOddRob = 0;
        for (int i = 0; i < nums.length; i++) {
            if ((i % 2) == 0) {
                mEvenRob = Math.max(mEvenRob + nums[i], mOddRob);
            } else {
                mOddRob = Math.max(mOddRob + nums[i], mEvenRob);
            }
        }

        return Math.max(mEvenRob, mOddRob);
    }

    // Use recursive with memorize
    // Time Complexity: O(n)
    // Space complexity: O(n)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 34.3 MB, less than 100.00%
    public int rob3(int[] nums) {
        Integer[] mMem = new Integer[nums.length];
        return rob(nums, nums.length - 1, mMem);
    }

    private int rob(int[] nums, int mIndex, Integer[] mMem) {
        if (mIndex < 0) {
            return 0;
        }
        if (mMem[mIndex] != null) {
            return mMem[mIndex];
        }

        mMem[mIndex] = Math.max(rob(nums, mIndex - 2, mMem) + nums[mIndex], rob(nums, mIndex - 1, mMem));
        return mMem[mIndex];
    }
}
