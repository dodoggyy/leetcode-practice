/**
 * 
 */
package com.medium;

import java.util.List;
import java.util.ArrayList;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class CombinationSumIII_216 {
    // 2019年12月18日
    // Use DFS
    // Time Complexity: O(2^9)
    // Space Complexity:O(n)
    // Runtime: 1 ms, faster than 55.43%
    // Memory Usage: 34 MB, less than 8.00%
    public List<List<Integer>> combinationSum3(int k, int n) {
        List<List<Integer>> mResult = new ArrayList<>();
        dfs(k, n, 0, 1, mResult, new ArrayList<Integer>());
        return mResult;
    }

    private void dfs(int k, int target, int sum, int mStart, List<List<Integer>> mResult, List<Integer> mCur) {
        if (sum == target && k == mCur.size()) {
            mResult.add(new ArrayList<>(mCur));
        }
        if (sum > target) {
            return;
        }

        for (int i = mStart; i <= 9; i++) {
            mCur.add(i);
            dfs(k, target, sum + i, i + 1, mResult, mCur);
            mCur.remove(mCur.size() - 1);
        }
    }
}
