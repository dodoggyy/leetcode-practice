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
public class InvertBinaryTree_226 {

    // Use Recursive:
    // Time Complexity: O(n)
    // Space Complexity: O(h) = O(n)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 35.7 MB, less than 29.71%
    public TreeNode invertTree(TreeNode root) {
        if (root == null) {
            return null;
        }
        TreeNode mLeftTree = root.left; // tmp left tree
        root.left = invertTree(root.right);
        root.right = invertTree(mLeftTree);

        return root;
    }

    // Use Iterative(BFS):
    // Time Complexity: O(n)
    // Space Complexity: O(n)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 35.7 MB, less than 29.71%
    public TreeNode invertTree2(TreeNode root) {
        if (root == null) {
            return root;
        }
        Queue<TreeNode> mQueue = new LinkedList<>();
        mQueue.add(root);

        while (!mQueue.isEmpty()) {
            TreeNode mNode = mQueue.poll();
            TreeNode mTmp = mNode.left;
            mNode.left = mNode.right;
            mNode.right = mTmp;
            if (mNode.left != null) {
                mQueue.add(mNode.left);
            }
            if (mNode.right != null) {
                mQueue.add(mNode.right);
            }
        }
        
        return root;
    }
}
