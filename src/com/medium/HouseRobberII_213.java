/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class HouseRobberII_213 {
    public static void main(String[] args) throws Exception {
        HouseRobberII_213 robber = new HouseRobberII_213();
        System.out.println(robber.rob(new int[] { 1, 2, 3, 4, 5, 1, 2, 3, 4, 5 }));
    }

    // consider 2 case:
    // case 1: rob 0 => rob(0:N-1)
    // case 2: not rob 0 => rob(1:N)
    public int rob(int[] nums) {
        if (nums.length == 0 || nums == null) {
            return 0;
        }
        if (nums.length == 1) {
            return nums[0];
        }
        // System.out.println(rob(nums, 0, nums.length - 2));
        // System.out.println(rob(nums, 1, nums.length - 1));
        return Math.max(rob2(nums, 0, nums.length - 2), rob2(nums, 1, nums.length - 1));
    }

    // Use DP with memorize
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 34 MB, less than 100.00%
    private int rob(int[] nums, int mIndexStart, int mIndexEnd) {
        int mLength = mIndexEnd - mIndexStart + 1;
        if (mLength == 0) {
            return 0;
        }
        if (mLength == 1) {
            return nums[mIndexStart];
        }
        int[] mDP = new int[mLength];
        mDP[0] = nums[mIndexStart];
        mDP[1] = Math.max(nums[mIndexStart], nums[mIndexStart + 1]);
        for (int i = 2; i < mLength; i++) {
            mDP[i] = Math.max(mDP[i - 2] + nums[mIndexStart + i], mDP[i - 1]);
            // System.out.printf("mDP[%d]:%d\n", i, mDP[i]);
        }
        return mDP[mLength - 1];
    }

    // Use DP without memorize
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 34.2 MB, less than 100.00%
    private int rob2(int[] nums, int mIndexStart, int mIndexEnd) {
        int mPrevious = 0, mCurrent = 0;
        for (int i = mIndexStart; i <= mIndexEnd; i++) {
            int mTmp = mCurrent;
            mCurrent = Math.max(mPrevious + nums[i], mCurrent);
            mPrevious = mTmp;
        }
        return mCurrent;
    }
}
