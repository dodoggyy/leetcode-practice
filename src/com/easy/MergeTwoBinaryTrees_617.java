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
public class MergeTwoBinaryTrees_617 {

    // Use recursive
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 1 ms, faster than 69.01%
    // Memory Usage: 40.1 MB, less than 100.00%
    public TreeNode mergeTrees(TreeNode t1, TreeNode t2) {
        if (t1 == null && t2 == null) {
            return null;
        }
        if (t1 == null) {
            return t2;
        }
        if (t2 == null) {
            return t1;
        }
        TreeNode mMergeTree = new TreeNode(t1.val + t2.val);
        mMergeTree.left = mergeTrees(t1.left, t2.left);
        mMergeTree.right = mergeTrees(t1.right, t2.right);

        return mMergeTree;
    }

    // Use iterative
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 1 ms, faster than 69.01%
    // Memory Usage: 39.7 MB, less than 100.00%
    public TreeNode mergeTrees2(TreeNode t1, TreeNode t2) {
        if (t1 == null) {
            return t2;
        }
        if (t2 == null) {
            return t1;
        }
        Queue<TreeNode[]> mQueue = new LinkedList<>();

        mQueue.add(new TreeNode[] { t1, t2 });

        while (!mQueue.isEmpty()) {
            TreeNode[] mNode = mQueue.poll();

            if (mNode[1] == null) {
                continue;
            }

            mNode[0].val += mNode[1].val;
            if (mNode[0].left == null) {
                mNode[0].left = mNode[1].left;
            } else {
                mQueue.add(new TreeNode[] { mNode[0].left, mNode[1].left });
            }
            if (mNode[0].right == null) {
                mNode[0].right = mNode[1].right;
            } else {
                mQueue.add(new TreeNode[] { mNode[0].right, mNode[1].right });
            }
        }
        return t1;
    }
}
