/**
 * 
 */
package com.easy;

import java.util.LinkedList;
import java.util.Queue;

/**
 * @author Chris Lin
 * @version 1.1
 */
public class MaximumDepthofBinaryTree_104 {

    // Use Recursive
    // Time Complexity: O(n)
    // Space Complexity: O(logn)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 38.8 MB, less than 94.62%
    public int maxDepth(TreeNode root) {
        if (root == null) {
            return 0;
        }
        // int mCountLeft = maxDepth(root.left);
        // int mCountRight = maxDepth(root.right);
        // return Math.max(mCountLeft, mCountRight) + 1;
        return Math.max(maxDepth(root.left), maxDepth(root.right)) + 1;

    }

    // Use Iterative
    // Time Complexity: O(n)
    // Space Complexity: O(n)
    // Runtime: 1 ms, faster than 11.59%
    // Memory Usage: 38.5 MB, less than 96.24%
    public int maxDepth2(TreeNode root) {
        int mResult = 0;
        if (root == null) {
            return mResult;
        }
        Queue<TreeNode> mQueue = new LinkedList<>();
        mQueue.add(root);
        while (!mQueue.isEmpty()) {
            int size = mQueue.size();
            for (int i = 0; i < size; i++) {
                TreeNode mNode = mQueue.poll();
                if (mNode.left != null) {
                    mQueue.add(mNode.left);
                }
                if (mNode.right != null) {
                    mQueue.add(mNode.right);
                }
            }
            mResult++;
        }

        return mResult;
    }
}
