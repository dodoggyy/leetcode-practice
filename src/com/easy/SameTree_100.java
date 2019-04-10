/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class SameTree_100 {
    public class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int x) {
            val = x;
        }
    }

    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 35.7 MB, less than 96.28%
    public boolean isSameTree(TreeNode p, TreeNode q) {
        // simple case
        if (p == null && q == null) {
            return true;
        }
        if ((p == null && q != null) || (p != null && q == null)) {
            return false;
        }

        return ((p.val == q.val) && isSameTree(p.left, q.left) && isSameTree(p.right, q.right));
    }
}
