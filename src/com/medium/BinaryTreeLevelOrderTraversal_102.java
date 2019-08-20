/**
 * 
 */
package com.medium;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class BinaryTreeLevelOrderTraversal_102 {
    // 2019年8月20日
    // Use BFS
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 1 ms, faster than 89.54%
    // Memory Usage: 36.2 MB, less than 100.00%
    public List<List<Integer>> levelOrder(TreeNode root) {
        List<List<Integer>> mResult = new ArrayList<>();

        if (root == null) {
            return mResult;
        }

        Queue<TreeNode> mQueue = new LinkedList<TreeNode>();
        mQueue.add(root);

        while (!mQueue.isEmpty()) {

            List<Integer> mList = new ArrayList<>();
            int size = mQueue.size();

            for (int i = 0; i < size; i++) {
                TreeNode mNode = mQueue.poll();
                if (mNode != null) {
                    mList.add(mNode.val);
                }
                if (mNode.left != null) {
                    mQueue.add(mNode.left);
                }

                if (mNode.right != null) {
                    mQueue.add(mNode.right);
                }
            }

            mResult.add(mList);

        }

        return mResult;
    }

    // 2019年8月20日
    // Use DFS (pre-order)
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 36.5 MB, less than 100.00%
    public List<List<Integer>> levelOrder2(TreeNode root) {
        List<List<Integer>> mResult = new ArrayList<>();

        DFS(root, 0, mResult);

        return mResult;
    }

    private void DFS(TreeNode root, int mDepth, List<List<Integer>> mResult) {
        if (root == null) {
            return;
        }
        if (mDepth >= mResult.size()) {
            mResult.add(new LinkedList<Integer>());
        }
        mResult.get(mDepth).add(root.val);
        DFS(root.left, mDepth + 1, mResult);
        DFS(root.right, mDepth + 1, mResult);
    }
}
