/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.1
 */
public class BalancedBinaryTree_110 {
    public boolean mResult = true;

    // Use recursive
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 82.21%
    // Memory Usage: 40.8 MB, less than 19.55%
    public boolean isBalanced(TreeNode root) {
        maxDepth(root);
        return mResult;
    }

    private int maxDepth(TreeNode root) {
        if (root == null) {
            return 0;
        }
        int mDepthLeft = maxDepth(root.left);
        int mDepthRight = maxDepth(root.right);

        if (Math.abs(mDepthLeft - mDepthRight) > 1) {
            mResult = false;
        }

        return Math.max(mDepthLeft, mDepthRight) + 1;
    }

    // Use recursive without global variable
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 99.42%
    // Memory Usage: 41.4 MB, less than 11.11%
    public boolean isBalanced2(TreeNode root) {
        if (root == null) {
            return true;
        }
        int mDiff = Math.abs(getDepth(root.left) - getDepth(root.right));
        if (mDiff <= 1) {
            return isBalanced2(root.left) && isBalanced2(root.right);
        } else {
            return false;
        }
    }

    private int getDepth(TreeNode root) {
        if (root == null) {
            return 0;
        }
        return Math.max(getDepth(root.left), getDepth(root.right)) + 1;
    }
}
