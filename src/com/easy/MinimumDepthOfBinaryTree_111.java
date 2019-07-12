/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class MinimumDepthOfBinaryTree_111 {

    // Use recursive:
    // Time Complexity: O(n)
    // Space Complexity: O(1)
    // Runtime: 0 ms, faster than 100.00% 
    // Memory Usage: 38.5 MB, less than 98.90%
    public int minDepth(TreeNode root) {
        if (root == null) {
            return 0;
        }
        int mLeft = minDepth(root.left);
        int mRight = minDepth(root.right);
        if (mLeft == 0 || mRight == 0) {
            return mLeft + mRight + 1;
        }

        return Math.min(mLeft, mRight) + 1;
    }
}
