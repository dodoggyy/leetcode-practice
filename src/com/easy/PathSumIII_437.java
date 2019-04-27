/**
 * 
 */
package com.easy;

import java.util.HashMap;
import java.util.Map;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class PathSumIII_437 {

    // Recursive version
    // Runtime: 11 ms, faster than 64.98%
    // Memory Usage: 42.1 MB, less than 11.25%
    public int pathSum(TreeNode root, int sum) {
        if (root == null) {
            return 0;
        }
        return pathSumStartWithRoot(root, sum) + pathSum(root.left, sum) + pathSum(root.right, sum);
    }

    private int pathSumStartWithRoot(TreeNode root, int sum) {
        if (root == null) {
            return 0;
        }
        int mResult = 0;
        if (root.val == sum) {
            mResult++;
        }
        mResult += pathSumStartWithRoot(root.left, sum - root.val) + pathSumStartWithRoot(root.right, sum - root.val);
        return mResult;
    }

    // DP version
    // Runtime: 3 ms, faster than 100.00%
    // Memory Usage: 40.2 MB, less than 69.00%
    public int pathSum2(TreeNode root, int sum) {
        if (root == null) {
            return 0;
        }
        Map<Integer, Integer> mMap = new HashMap<>();
        mMap.put(0, 1); // initial value
        return pathSumStartWithRoot(root, sum, mMap, 0);
    }

    private int pathSumStartWithRoot(TreeNode root, int sum, Map<Integer, Integer> mMap, int mPreviousDP) {
        if (root == null) {
            return 0;
        }
        int mCurrent = mPreviousDP + root.val;
        int mResult = mMap.getOrDefault(mCurrent - sum, 0);
        mMap.put(mCurrent, mMap.getOrDefault(mCurrent, 0) + 1);
        mResult += pathSumStartWithRoot(root.left, sum, mMap, mCurrent)
                + pathSumStartWithRoot(root.right, sum, mMap, mCurrent);
        mMap.put(mCurrent, mMap.get(mCurrent) - 1); // update to hash map
        return mResult;
    }
}
