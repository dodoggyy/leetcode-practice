/**
 * 
 */
package com.medium;

import java.util.HashMap;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class HouseRobberIII_337 {

    // Compare between (grandparnet + MAX(grandchildren)) v.s. (MAX(children))
    // without memorize version
    // Time Complexity: O(n^2)
    // Space Complexity: O(n)
    // Runtime: 610 ms, faster than 17.98%
    // Memory Usage: 38.2 MB, less than 90.53%
    public int rob(TreeNode root) {
        if (root == null) {
            return 0;
        }
        int mValue = root.val;
        int mLL = (root.left != null) ? rob(root.left.left) : 0;
        int mLR = (root.left != null) ? rob(root.left.right) : 0;
        int mRL = (root.right != null) ? rob(root.right.left) : 0;
        int mRR = (root.right != null) ? rob(root.right.right) : 0;

        return Math.max(mValue + mLL + mLR + mRL + mRR, rob(root.left) + rob(root.right));

    }

    // with memorize version
    // Time Complexity: O(n)
    // Space Complexity: O(n)
    // Runtime: 4 ms, faster than 53.26%
    // Memory Usage: 42.6 MB, less than 8.33%
    private HashMap<TreeNode, Integer> mMap = new HashMap<>();

    public int rob2(TreeNode root) {
        if (root == null) {
            return 0;
        }
        if (mMap.containsKey(root)) {
            return mMap.get(root);
        }
        int mValue = root.val;
        int mLL = (root.left != null) ? rob(root.left.left) : 0;
        int mLR = (root.left != null) ? rob(root.left.right) : 0;
        int mRL = (root.right != null) ? rob(root.right.left) : 0;
        int mRR = (root.right != null) ? rob(root.right.right) : 0;
        mMap.put(root, Math.max(mValue + mLL + mLR + mRL + mRR, rob(root.left) + rob(root.right)));

        return mMap.get(root);
    }

    // Use greedy
    // Time Complexity: O(n)
    // Space Complexity: O(1)
    // Runtime: 1 ms, faster than 99.17%
    // Memory Usage: 38.7 MB, less than 72.22%
    public int rob3(TreeNode root) {
        int[] mResult = robSub(root);
        return Math.max(mResult[0], mResult[1]);
    }

    private int[] robSub(TreeNode root) {
        int[] mResult = new int[2];

        if (root == null) {
            return mResult;
        }

        int[] mLeft = robSub(root.left);
        int[] mRight = robSub(root.right);

        mResult[0] = root.val + mLeft[1] + mRight[1];
        mResult[1] = Math.max(mLeft[0], mLeft[1]) + Math.max(mRight[0], mRight[1]);

        return mResult;
    }
}
