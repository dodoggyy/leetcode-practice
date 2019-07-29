/**
 * 
 */
package com.easy;

import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class MaximumDepthOfNAryTree_559 {
    class Node {
        public int val;
        public List<Node> children;

        public Node() {
        }

        public Node(int _val, List<Node> _children) {
            val = _val;
            children = _children;
        }
    };

    // 2019年7月29日
    // Use DFS
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 1 ms, faster than 98.24%
    // Memory Usage: 44.6 MB, less than 84.85%
    public int maxDepth(Node root) {
        if (root == null) {
            return 0;
        }
        int mDepth = 0;
        for (Node child : root.children) {
            mDepth = Math.max(mDepth, maxDepth(child));
        }
        return mDepth + 1;
    }

    // 2019年7月29日
    // Use BFS
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 2 ms, faster than 28.12%
    // Memory Usage: 48.3 MB, less than 36.42%
    public int maxDepth2(Node root) {
        if (root == null) {
            return 0;
        }

        Queue<Node> mQueue = new LinkedList<>();
        mQueue.add(root);

        int mDepth = 0;

        while (!mQueue.isEmpty()) {
            int mSize = mQueue.size();
            for (int i = 0; i < mSize; i++) {
                Node mCurrent = mQueue.poll();
                for (Node child : mCurrent.children) {
                    mQueue.add(child);
                }
            }
            mDepth++;
        }

        return mDepth;
    }
}
