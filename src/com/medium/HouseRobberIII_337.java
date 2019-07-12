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
}
