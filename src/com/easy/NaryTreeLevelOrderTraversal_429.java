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
public class NaryTreeLevelOrderTraversal_429 {
    // 2019年7月19日
    // Use BFS
    // Time Complexity: O(n)
    // Space Complexity:O(h)
    // Runtime: 3 ms, faster than 75.82%
    // Memory Usage: 46.6 MB, less than 82.51%
    public List<List<Integer>> levelOrder(Node root) {
        List<List<Integer>> mResult = new ArrayList<>();
        if (root == null) {
            return mResult;
        }
        Queue<Node> mQueue = new LinkedList<>();
        mQueue.add(root);

        while (!mQueue.isEmpty()) {
            List<Integer> mList = new ArrayList<>();
            int size = mQueue.size();
            for (int i = 0; i < size; i++) {
                Node mNode = mQueue.poll();
                mList.add(mNode.val);
                if (mNode.children != null) {
                    for (Node node : mNode.children) {
                        mQueue.add(node);
                    }

                }
            }
            mResult.add(mList);
        }

        return mResult;
    }

    // 2019年7月19日
    // Use DFS
    // Time Complexity: O(n)
    // Space Complexity:O(h)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 46.4 MB, less than 84.19%
    private List<List<Integer>> mResult2 = new ArrayList<>();

    public List<List<Integer>> levelOrder2(Node root) {
        preorder(root, 1);

        return mResult2;
    }

    private void preorder(Node root, int mDepth) {
        if (root == null) {
            return;
        }
        if (mResult2.size() < mDepth) {
            mResult2.add(new ArrayList<>());
        }
        mResult2.get(mDepth - 1).add(root.val);
        for (Node child : root.children) {
            preorder(child, mDepth + 1);
        }
    }

}
