/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class LongestUnivaluePath_687 {

    // Use recursive:
    // Time Complexity: O(n)
    // Space Complexity: O(n)
    // Runtime: 4 ms, faster than 65.47%
    // Memory Usage: 40.6 MB, less than 98.87%
    private int mAnswer = 0;

    public int longestUnivaluePath(TreeNode root) {
        if (root == null) {
            return 0;
        }
        univaluePath(root);

        return mAnswer;
    }

    public int univaluePath(TreeNode root) {
        if (root == null) {
            return 0;
        }
        int mLeft = univaluePath(root.left);
        int mRight = univaluePath(root.right);
        int mPathLeft = 0, mPathRight = 0;
        if (root.left != null && root.val == root.left.val) {
            mPathLeft = mLeft + 1;
        }
        if (root.right != null && root.val == root.right.val) {
            mPathRight = mRight + 1;
        }
        // compare and store longer path
        mAnswer = Math.max(mAnswer, mPathLeft + mPathRight);

        // only return one side maximum in recursive
        return Math.max(mPathLeft, mPathRight);
    }
}
