/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.1
 */
public class DiameterofBinaryTree_543 {
    public int mResult = 0;

    // Use Recursive
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 37.7 MB, less than 94.26%
    public int diameterOfBinaryTree(TreeNode root) {
        maxDepth(root);

        return mResult;
    }

    private int maxDepth(TreeNode root) {
        if (root == null) {
            return 0;
        }
        int mDepthLeft = maxDepth(root.left);
        int mDepthRight = maxDepth(root.right);
        mResult = Math.max(mDepthLeft + mDepthRight, mResult);

        return Math.max(mDepthLeft, mDepthRight) + 1;
    }

    // Use Recursive without global variable
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 10 ms, faster than 7.69%
    // Memory Usage: 38.5 MB, less than 36.36% 
    public int diameterOfBinaryTree2(TreeNode root) {
        if (root == null) {
            return 0;
        }
        int mDepthLeft = getDepth(root.left);
        int mDepthRight = getDepth(root.right);

        int mDiameterLeft = diameterOfBinaryTree2(root.left);
        int mDiameterRight = diameterOfBinaryTree2(root.right);

        return Math.max(mDepthLeft + mDepthRight, Math.max(mDiameterLeft, mDiameterRight));
    }

    private int getDepth(TreeNode root) {
        return (root == null) ? 0 : Math.max(getDepth(root.left), getDepth(root.right)) + 1;
    }
}
