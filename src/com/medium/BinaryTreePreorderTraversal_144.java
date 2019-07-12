/**
 * 
 */
package com.medium;

import java.util.ArrayList;
import java.util.List;
import java.util.Stack;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class BinaryTreePreorderTraversal_144 {
    // Recursive version (DLR)
    // Time Complexity: O(n)
    // Space Complexity: O(n)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 34.7 MB, less than 99.98%
    private List<Integer> mAnswer = new ArrayList<>();

    public List<Integer> preorderTraversal(TreeNode root) {
        if (root == null) {
            return mAnswer;
        }
        mAnswer.add(root.val);
        preorderTraversal(root.left);
        preorderTraversal(root.right);

        return mAnswer;
    }

    // Iterative version (DLR)
    // Time Complexity: O(n)
    // Space Complexity: O(n)
    // Runtime: 1 ms, faster than 56.33%
    // Memory Usage: 34.8 MB, less than 99.98%
    public List<Integer> preorderTraversal2(TreeNode root) {
        Stack<TreeNode> mStack = new Stack<>();
        mStack.push(root);
        while (!mStack.isEmpty()) {
            TreeNode mNode = mStack.pop();
            if (mNode == null) {
                continue;
            }
            mAnswer.add(mNode.val);
            mStack.push(mNode.right);
            mStack.push(mNode.left);
        }

        return mAnswer;
    }

}
