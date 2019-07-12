/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class LowestCommonAncestorOfABinaryTree_236 {
    // Use recursive
    // Time Complexity: O(n)
    // Space Complexity: O(h)
    // Runtime: 6 ms, faster than 55.17%
    // Memory Usage: 35.3 MB, less than 5.01%
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if (root == null || p == root || q == root) {
            return root;
        }
        TreeNode mLeftTree = lowestCommonAncestor(root.left, p, q);
        TreeNode mRightTree = lowestCommonAncestor(root.right, p, q);
        if (mLeftTree != null && mRightTree != null) {
            return root;
        } else {
            return (mLeftTree == null) ? mRightTree : mLeftTree;
        }
    }
}
