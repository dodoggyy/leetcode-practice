/**
 * 
 */
package com.easy;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

/**
 * @author Chris Lin
 * @version 1.0
 * 
 */
public class BinaryTreePaths_257 {

    private List<String> mResult = new ArrayList<String>();

    // Use DFS
    // Time Complexity: O(n)
    // Space Complexity: O(n)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 36.5 MB, less than 100.00%
    public List<String> binaryTreePaths(TreeNode root) {
        if (root == null) {
            return mResult;
        }
        binaryTreePaths(root, root.val + "");

        return mResult;
    }

    private void binaryTreePaths(TreeNode root, String mPath) {
        if (root.left == null && root.right == null) {
            mResult.add(mPath);
        }
        if (root.left != null) {
            binaryTreePaths(root.left, mPath + "->" + root.left.val);
        }
        if (root.right != null) {
            binaryTreePaths(root.right, mPath + "->" + root.right.val);
        }
    }

    // Use BFS
    // Time Complexity: O(n)
    // Space Complexity: O(n)
    // Runtime: 2 ms, faster than 19.65%
    // Memory Usage: 36.2 MB, less than 100.00%
    public List<String> binaryTreePaths2(TreeNode root) {
        if (root == null) {
            return mResult;
        }

        Queue<TreeNode> mNodeQueue = new LinkedList<>();
        Queue<String> mPathQueue = new LinkedList<>();
        mNodeQueue.add(root);
        mPathQueue.add(root.val + "");

        while (!mNodeQueue.isEmpty()) {
            TreeNode mNode = mNodeQueue.poll();
            String mSinglePath = mPathQueue.poll();
            if (mNode.left == null && mNode.right == null) {
                mResult.add(mSinglePath);
            }
            if (mNode.left != null) {
                mNodeQueue.add(mNode.left);
                mPathQueue.add(mSinglePath + "->" + mNode.left.val);
            }
            if (mNode.right != null) {
                mNodeQueue.add(mNode.right);
                mPathQueue.add(mSinglePath + "->" + mNode.right.val);
            }
        }

        return mResult;
    }

}
