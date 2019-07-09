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
 */
public class BinaryTreeLevelOrderTraversalII_107 {
    // Use BFS traversal
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 1 ms, faster than 95.62%
    // Memory Usage: 36.1 MB, less than 99.91%
    private List<List<Integer>> mResult = new ArrayList<>();

    public List<List<Integer>> levelOrderBottom(TreeNode root) {
        if (root == null)
            return mResult;
        leverOrder(root);

        return mResult;
    }

    private void leverOrder(TreeNode root) {
        Queue<TreeNode> mQueue = new LinkedList<>();
        List<Integer> mList = new ArrayList<>();
        mQueue.add(root);
        while (!mQueue.isEmpty()) {
            mList = new ArrayList<>();
            int mSize = mQueue.size();
            for (int i = 0; i < mSize; i++) {
                TreeNode mNode = mQueue.poll();
                if (mNode != null) {
                    mList.add(mNode.val);
                    mQueue.add(mNode.left);
                    mQueue.add(mNode.right);
                }
            }
            if (!mList.isEmpty()) {
                mResult.add(0, mList);
            }
        }
    }
}
