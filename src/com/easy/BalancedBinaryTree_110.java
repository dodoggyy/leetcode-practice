/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class BalancedBinaryTree_110 {
    public boolean mResult = true;
   
    // Runtime: 1 ms, faster than 82.21%
    // Memory Usage: 40.8 MB, less than 19.55%
    public boolean isBalanced(TreeNode root) {
        maxDepth(root);
        return mResult;
    }
    
    private int maxDepth(TreeNode root) {
        if(root == null) {
            return 0;
        }
        int mDepthLeft = maxDepth(root.left);
        int mDepthRight = maxDepth(root.right);
        
        if(Math.abs(mDepthLeft - mDepthRight) > 1) {
            mResult = false;
        }
        
        return Math.max(mDepthLeft, mDepthRight) + 1;
    }
}
