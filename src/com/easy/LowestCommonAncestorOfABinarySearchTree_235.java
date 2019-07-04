/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class LowestCommonAncestorOfABinarySearchTree_235 {
    // Use recursive
    // Time Complexity: O(n)
    // Space Complexity: O(n)
    // Runtime: 4 ms, faster than 100.00%
    // Memory Usage: 36.6 MB, less than 5.04%
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if ((root.val < p.val) && (root.val < q.val)) {
            return lowestCommonAncestor(root.right, p, q);
        }
        if ((root.val > p.val) && (root.val > q.val)) {
            return lowestCommonAncestor(root.left, p, q);
        }

        return root;
    }
}
