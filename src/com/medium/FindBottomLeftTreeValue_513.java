/**
 * 
 */
package com.medium;

import java.util.LinkedList;
import java.util.Queue;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class FindBottomLeftTreeValue_513 {
    // Use traversal DRL
    // Time Complexity: O(n)
    // Space Complexity: O(n)
    // Runtime: 1 ms, faster than 69.84%
    // Memory Usage: 39 MB, less than 86.20%
    private TreeNode mNode;

    public int findBottomLeftValue(TreeNode root) {
        if (root == null) {
            return 0;
        }
        Queue<TreeNode> mQueue = new LinkedList<>();
        mQueue.add(root);

        // D -> R -> L, then last is the leftmost node
        while (mQueue.isEmpty() == false) {
            mNode = mQueue.poll();

            if (mNode.right != null) {
                mQueue.add(mNode.right);
            }
            if (mNode.left != null) {
                mQueue.add(mNode.left);
            }
        }
        return mNode.val;

    }
}
