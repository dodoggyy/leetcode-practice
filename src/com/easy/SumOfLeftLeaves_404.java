/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class SumOfLeftLeaves_404 {
    // Use recursive:
    // Time Complexity: O(n)
    // Space Complexity: O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 36.4 MB, less than 100.00%
    public int sumOfLeftLeaves(TreeNode root) {
        if (root == null) {
            return 0;
        }
        if(bIsLeaf(root.left)) {
            return root.left.val + sumOfLeftLeaves(root.right);
        }
        return (sumOfLeftLeaves(root.left) + sumOfLeftLeaves(root.right));
    }

    public boolean bIsLeaf(TreeNode node) {
        if (node == null) {
            return false;
        }
        return ((node.left == null) && (node.right == null));
    }
}
