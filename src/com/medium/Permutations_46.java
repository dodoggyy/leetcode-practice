/**
 * 
 */
package com.medium;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class Permutations_46 {
    // 2019年8月10日
    // Use backtrack
    // Time Complexity: O(n^2)
    // Space Complexity:O(n!)
    // Runtime: 1 ms, faster than 98.21%
    // Memory Usage: 37.7 MB, less than 90.07%
    private List<List<Integer>> mResult = new ArrayList<>();

    public List<List<Integer>> permute(int[] nums) {
        if (nums == null || nums.length == 0) {
            return mResult;
        }

        backTracking(new ArrayList<>(), nums);

        return mResult;
    }

    private void backTracking(List<Integer> mTmp, int[] nums) {
        if (mTmp.size() == nums.length) {
            mResult.add(new ArrayList<>(mTmp));
        } else {
            for (int i = 0; i < nums.length; i++) {
                if (!mTmp.contains(nums[i])) {
                    mTmp.add(nums[i]);
                    backTracking(mTmp, nums);
                    // remove last element to recover
                    // ex: [1,2,3]
                    // remove order: (3 2 2 3 1), (3 1 1 3 2),  (2 1 1 2 3)
                    // Round 1: 1,2,3 => 1,2 => 1 => 1,3 => 1,3,2 => 1,3 => 1 =>
                    // Round 2: 2,1,3 => 2,1 => 2 => 2,3,1 => 2,3 => 2 =>
                    // Round 3: 3,1,2 => 3,1 => 3 => 3,2,1 => 3,2 => 3 =>
                    mTmp.remove(mTmp.size() - 1);
                }
            }
        }
    }
}
