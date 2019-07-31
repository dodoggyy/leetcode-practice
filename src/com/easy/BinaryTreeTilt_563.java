/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class BinaryTreeTilt_563 {
    // 2019年7月31日
    // Use post order
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 1 ms, faster than 69.51%
    // Memory Usage: 39.3 MB, less than 79.17%
    int mResult = 0;

    public int findTilt(TreeNode root) {
        postorder(root);
        return mResult;
    }

    private int postorder(TreeNode root) {
        if (root == null) {
            return 0;
        }

        int mLeft = postorder(root.left);
        int mRight = postorder(root.right);

        mResult += Math.abs(mLeft - mRight);

        return mLeft + mRight + root.val;
    }

    // 2019年7月31日
    // Use recursive
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 11 ms, faster than 5.45%
    // Memory Usage: 37.9 MB, less than 99.40%
    public int findTilt2(TreeNode root) {
        if (root == null) {
            return 0;
        }
        return Math.abs(countValue(root.left) - countValue(root.right)) + findTilt2(root.left) + findTilt2(root.right);
    }

    private int countValue(TreeNode root) {
        if (root == null) {
            return 0;
        }
        return root.val + countValue(root.left) + countValue(root.right);
    }
}
