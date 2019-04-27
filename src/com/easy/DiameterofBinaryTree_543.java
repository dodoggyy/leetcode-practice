/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class DiameterofBinaryTree_543 {
    public int mResult = 0;
    
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 37.7 MB, less than 94.26%
    public int diameterOfBinaryTree(TreeNode root) {
        maxDepth(root);
        
        return mResult;
    }
    
    private int maxDepth(TreeNode root) {
        if(root == null) {
            return 0;
        }
        int mDepthLeft = maxDepth(root.left);
        int mDepthRight = maxDepth(root.right);
        mResult = Math.max(mDepthLeft + mDepthRight, mResult);

        return Math.max(mDepthLeft, mDepthRight) + 1;
    }
}
