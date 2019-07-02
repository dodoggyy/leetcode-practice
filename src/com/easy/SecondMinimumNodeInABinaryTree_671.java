/**
 * 
 */
package com.easy;

import java.util.Deque;
import java.util.LinkedList;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class SecondMinimumNodeInABinaryTree_671 {

    public int findSecondMinimumValue(TreeNode root) {
        if (root == null) {
            return -1;
        }
        // return BFS(root);
        return DFS(root, root.val);
    }

    // Use DFS:
    // Time Complexity: O(n)
    // Space Complexity: O(n)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 34.2 MB, less than 100.00%
    private int DFS(TreeNode root, int mMin) {
        if (root == null) {
            return -1;
        }

        // if root's value greater than mMin,
        // then according to tree's characteristic all children's value should
        // be >= mMin
        // So no need to visit children
        if (root.val > mMin) {
            return root.val;
        }

        int mMinLeft = DFS(root.left, mMin);
        int mMinRight = DFS(root.right, mMin);

        if (mMinLeft == -1) {
            return mMinRight;
        }
        if (mMinRight == -1) {
            return mMinLeft;
        }

        return Math.min(mMinLeft, mMinRight);
    }

    // Use BFS:
    // Time Complexity: O(n)
    // Space Complexity: O(n)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 33.9 MB, less than 100.00%
    private int BFS(TreeNode root) {
        if (root == null) {
            return -1;
        }

        int mMin = root.val;
        int mSecondMin = Integer.MAX_VALUE;
        boolean bIsFound = false;

        Deque<TreeNode> mDeQueue = new LinkedList<>();
        mDeQueue.push(root);

        while (!mDeQueue.isEmpty()) {
            TreeNode mNode = mDeQueue.pop();

            // consider value may equal integer max
            if (mNode.val > mMin && mNode.val <= mSecondMin) {
                mSecondMin = mNode.val;
                bIsFound = true;
                continue;
            }

            if (mNode.left == null) {
                continue;
            }
            mDeQueue.push(mNode.left);
            mDeQueue.push(mNode.right);
        }

        return bIsFound ? mSecondMin : -1;
    }
}
