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
 * 
 */
public class NaryTreePreorderTraversal_589 {
    // Use Recursive
    // Time Complexity: O(n)
    // Space Complexity: O(n)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 44.4 MB, less than 96.54%
    private List<Integer> mResult = new ArrayList<>();

    public List<Integer> preorder(Node root) {
        helper(root);
        return mResult;

    }

    private void helper(Node root) {
        if (root == null) {
            return;
        }
        mResult.add(root.val);
        for (int i = 0; i < root.children.size(); i++) {
            helper(root.children.get(i));
        }
    }

    // Use Iterative
    // Time Complexity: O(n)
    // Space Complexity: O(n)
    // Runtime: 3 ms, faster than 36.47%
    // Memory Usage: 49.3 MB, less than 17.65%
    public List<Integer> preorder2(Node root) {
        if (root == null) {
            return new ArrayList<>();
        }
        List<Integer> mResult = new ArrayList<>();
        Stack<Node> mStack = new Stack<>();
        mStack.add(root);

        while (!mStack.isEmpty()) {
            Node mNode = mStack.pop();
            mResult.add(mNode.val);

            for (int i = mNode.children.size() - 1; i >= 0; i--) {
                mStack.add(mNode.children.get(i));
            }
        }

        return mResult;
    }
}
