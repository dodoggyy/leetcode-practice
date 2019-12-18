/**
 * 
 */
package com.medium;

import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class CombinationSumII_40 {
    // 2019年12月18日
    // Use DFS
    // Time Complexity: O(n^d)
    // Space Complexity:O(n)
    // n for candidate nums which less than target
    // d for depth that d = target/ min
    // Runtime: 25 ms, faster than 9.00%
    // Memory Usage: 38.5 MB, less than 88.42%
    public List<List<Integer>> combinationSum2(int[] candidates, int target) {
        List<List<Integer>> mResult = new ArrayList<>();
        Arrays.sort(candidates);
        dfs(candidates, target, 0, 0, mResult, new ArrayList<Integer>());
        return mResult;
    }

    private void dfs(int[] candidates, int target, int sum, int mStart, List<List<Integer>> mResult,
            List<Integer> mCur) {
        if (sum == target && !mResult.contains(mCur)) {
            mResult.add(new ArrayList<>(mCur));
        }
        if (sum > target) {
            return;
        }

        for (int i = mStart; i < candidates.length; i++) {
            mCur.add(candidates[i]);
            dfs(candidates, target, sum + candidates[i], i + 1, mResult, mCur);
            mCur.remove(mCur.size() - 1);
        }

    }

    // 2019年12月18日
    // Optimize DFS
    // Time Complexity: O(n^d)
    // Space Complexity:O(n)
    // n for candidate nums which less than target
    // d for depth that d = target/ min
    // Runtime: 7 ms, faster than 28.59%
    // Memory Usage: 37.9 MB, less than 93.68%
    public List<List<Integer>> combinationSum3(int[] candidates, int target) {
        List<List<Integer>> mResult = new ArrayList<>();
        Arrays.sort(candidates);
        dfs2(candidates, target, 0, 0, mResult, new ArrayList<Integer>());
        return mResult;
    }

    private void dfs2(int[] candidates, int target, int sum, int mStart, List<List<Integer>> mResult,
            List<Integer> mCur) {
        if (sum == target) {
            mResult.add(new ArrayList<>(mCur));
        }
        if (sum > target) {
            return;
        }

        for (int i = mStart; i < candidates.length; i++) {
            if (i != mStart && candidates[i] == candidates[i - 1]) {
                continue;
            }
            mCur.add(candidates[i]);
            dfs2(candidates, target, sum + candidates[i], i + 1, mResult, mCur);
            mCur.remove(mCur.size() - 1);
        }

    }
}
