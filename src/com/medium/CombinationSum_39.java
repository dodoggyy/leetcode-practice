/**
 * 
 */
package com.medium;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class CombinationSum_39 {
    // 2019年12月18日
    // Use DFS
    // Time Complexity: O(n^d)
    // Space Complexity:O(n)
    // n for candidate nums which less than target
    // d for depth that d = target/ min
    // Runtime: 5 ms, faster than 52.26%
    // Memory Usage: 37.6 MB, less than 100.00%
    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        List<List<Integer>> mResult = new ArrayList<>();
        Arrays.sort(candidates);
        dfs(candidates, target, 0, 0, new ArrayList<Integer>(), mResult);

        return mResult;
    }

    private void dfs(int[] candidates, int target, int sum, int mStart, List<Integer> mCur,
            List<List<Integer>> mResult) {
        if (sum > target) {
            return;
        }
        if (sum == target) {
            mResult.add(new ArrayList<Integer>(mCur));
            return;
        }

        for (int i = mStart; i < candidates.length; i++) {
            mCur.add(candidates[i]);
            dfs(candidates, target, sum + candidates[i], i, mCur, mResult);
            mCur.remove(mCur.size() - 1);
        }
    }
}
