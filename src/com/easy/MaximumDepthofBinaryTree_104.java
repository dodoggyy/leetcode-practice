/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class MaximumDepthofBinaryTree_104 {
    
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 40.3 MB, less than 5.01%
    public int maxDepth(TreeNode root) {
        if (root == null) {
            return 0;
        } else {
//            int mCountLeft = maxDepth(root.left);
//            int mCountRight = maxDepth(root.right);
//            return Math.max(mCountLeft, mCountRight) + 1;
            return Math.max(maxDepth(root.left), maxDepth(root.right)) + 1;
        }
    }
}
