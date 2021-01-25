/**
 * 
 */
package com.easy;

import java.util.ArrayList;
import java.util.List;
import java.util.Stack;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class NaryTreePostorderTraversal_590 {
    // 2019年7月19日
    // Use Recursive
    // Time Complexity: O(n)
    // Space Complexity:O(h)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 44.4 MB, less than 95.79%
    private List<Integer> mResult = new ArrayList<>();

    public List<Integer> postorder(Node root) {
        helper(root);
        return mResult;
    }

    private void helper(Node root) {
        if (root == null) {
            return;
        }
        for (Node child : root.children) {
            helper(child);
        }
        mResult.add(root.val);
    }

    // 2019年7月19日
    // Use Iterative
    // Time Complexity: O(n)
    // Space Complexity:O(h)
    // Runtime: 5 ms, faster than 15.42%
    // Memory Usage: 48.8 MB, less than 18.78%

    public List<Integer> postorder2(Node root) {
        List<Integer> mResult = new ArrayList<Integer>();
        if (root == null) {
            return mResult;
        }
        Stack<Node> mStack = new Stack<>();
        mStack.add(root);

        while (!mStack.isEmpty()) {
            Node mNode = mStack.pop();
            mResult.add(0, mNode.val);
            if (mNode.children != null) {
                for (Node child : mNode.children) {
                    mStack.add(child);
                }
            }

        }
        return mResult;
    }
}
