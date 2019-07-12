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
public class AverageOfLevelsInBinaryTree_637 {

    // Use BFS
    // Time Complexity: O(n)
    // Space Complexity: O(h)
    // Runtime: 3 ms, faster than 25.77%
    // Memory Usage: 41.7 MB, less than 42.96%
    public List<Double> averageOfLevels(TreeNode root) {
        List<Double> mAnswer = new ArrayList<>();
        Queue<TreeNode> mQueueCurrent = new LinkedList<>();
        Queue<TreeNode> mQueueNext = new LinkedList<>();
        mQueueCurrent.add(root);

        while (!mQueueCurrent.isEmpty()) {
            double mSum = 0;
            for (TreeNode mNode : mQueueCurrent) {
                mSum += mNode.val;
                if (mNode.left != null) {
                    mQueueNext.add(mNode.left);
                }
                if (mNode.right != null) {
                    mQueueNext.add(mNode.right);
                }
            }
            mAnswer.add(mSum / mQueueCurrent.size());

            mQueueCurrent = mQueueNext;
            mQueueNext = new LinkedList<>();

        }

        return mAnswer;
    }

}
