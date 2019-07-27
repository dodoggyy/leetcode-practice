/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class MinimumAbsoluteDifferenceInBST_530 {
    // 2019年7月27日
    // Use inorder traversal and comparison with sorted list
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 96.24%
    // Memory Usage: 39.1 MB, less than 91.91%
    private int mResult = Integer.MAX_VALUE;
    private int mPrevious = Integer.MIN_VALUE;

    public int getMinimumDifference(TreeNode root) {
        inorder(root);
        return mResult;
    }

    private void inorder(TreeNode root) {
        if (root == null) {
            return;
        }

        inorder(root.left);
        if (mPrevious != Integer.MIN_VALUE) {
            mResult = Math.min(mResult, root.val - mPrevious);
        }

        mPrevious = root.val;
        inorder(root.right);
    }
}
