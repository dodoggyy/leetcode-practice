/**
 * 
 */
package com.easy;

import java.util.Arrays;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class RelativeRanks_506 {
    // 2019年7月24日
    // Use index sorting with 1 dimension array
    // Time Complexity: O(nlogn)
    // Space Complexity:O(n)
    // Runtime: 38 ms, faster than 25.39%
    // Memory Usage: 38.1 MB, less than 96.45%
    public String[] findRelativeRanks(int[] nums) {
        Integer[] mIndex = new Integer[nums.length];

        for (int i = 0; i < nums.length; i++) {
            mIndex[i] = i;
        }

        Arrays.sort(mIndex, (a, b) -> (nums[b] - nums[a]));

        String[] mResult = new String[nums.length];

        for (int i = 0; i < nums.length; i++) {
            if (i == 0) {
                mResult[mIndex[i]] = "Gold Medal";
            } else if (i == 1) {
                mResult[mIndex[i]] = "Silver Medal";
            } else if (i == 2) {
                mResult[mIndex[i]] = "Bronze Medal";
            } else {
                mResult[mIndex[i]] = (i + 1) + "";
            }
        }

        return mResult;
    }

    // 2019年7月24日
    // Use index sorting with 2 dimension array
    // Time Complexity: O(nlogn)
    // Space Complexity:O(n)
    // Runtime: 38 ms, faster than 25.39%
    // Memory Usage: 37.6 MB, less than 98.58%
    public String[] findRelativeRanks2(int[] nums) {
        Integer[][] mPair = new Integer[nums.length][2];

        for (int i = 0; i < nums.length; i++) {
            mPair[i][0] = nums[i];
            mPair[i][1] = i;
        }

        Arrays.sort(mPair, (a, b) -> b[0] - a[0]);

        String[] mResult = new String[nums.length];

        for (int i = 0; i < nums.length; i++) {
            if (i == 0) {
                mResult[mPair[i][1]] = "Gold Medal";
            } else if (i == 1) {
                mResult[mPair[i][1]] = "Silver Medal";
            } else if (i == 2) {
                mResult[mPair[i][1]] = "Bronze Medal";
            } else {
                mResult[mPair[i][1]] = (i + 1) + "";
            }
        }

        return mResult;
    }
}
