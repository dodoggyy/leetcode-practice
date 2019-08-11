/**
 * 
 */
package com.medium;

import java.util.ArrayList;
import java.util.List;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class Subsets_78 {
    // 2019年8月11日
    // Use backtrack
    // Time Complexity: O(n*2^n)
    // Space Complexity:O(n)
    // Runtime: 1 ms, faster than 44.44%
    // Memory Usage: 37.2 MB, less than 99.18%
    private List<List<Integer>> mResult = new ArrayList<List<Integer>>();

    public List<List<Integer>> subsets(int[] nums) {
        backTracking(nums, new ArrayList<>(), 0);

        return mResult;
    }

    private void backTracking(int[] nums, List<Integer> mTmp, int mStart) {
        mResult.add(new ArrayList<>(mTmp));

        for (int i = mStart; i < nums.length; i++) {
            if (!mTmp.contains(nums[i])) {
                mTmp.add(nums[i]);
                backTracking(nums, mTmp, i + 1);
                mTmp.remove(mTmp.size() - 1);
            }
        }
    }

    // 2019年8月11日
    // Use bit operation
    // Time Complexity: O(n*2^n)
    // Space Complexity:O(n)
    // Runtime: 1 ms, faster than 44.44%
    // Memory Usage: 37.2 MB, less than 99.18%
    public List<List<Integer>> subsets2(int[] nums) {
        for (int i = 0; i < Math.pow(2, nums.length); i++) {

            List<Integer> mTmp = new ArrayList<>();
            for (int j = 0; j < nums.length; j++) {
                if (((1 << j) & i) != 0) { // 0:not take, 1:take
                    mTmp.add(nums[j]);
                }
            }
            mResult.add(mTmp);
        }

        return mResult;
    }
}
