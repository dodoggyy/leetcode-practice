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
public class BinaryTreeInorderTraversal_94 {

    private List<Integer> mAnswer = new ArrayList<>();

    // Recursive version (LDR)
    // Time Complexity: O(n)
    // Space Complexity: O(h)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 35 MB, less than 99.99%
    public List<Integer> inorderTraversal(TreeNode root) {
        if (root == null) {
            return mAnswer;
        }
        inorderTraversal(root.left);
        mAnswer.add(root.val);
        inorderTraversal(root.right);

        return mAnswer;
    }

    // Iterative version (LDR)
    // Time Complexity: O(n)
    // Space Complexity: O(h)
    // Runtime: 1 ms, faster than 53.57%
    // Memory Usage: 34.7 MB, less than 99.99%
    public List<Integer> inorderTraversal2(TreeNode root) {
        if (root == null) {
            return mAnswer;
        }
        Stack<TreeNode> mStack = new Stack<>();
        TreeNode mCurrent = root;
        while (mCurrent != null || !mStack.isEmpty()) {
            while (mCurrent != null) {
                mStack.push(mCurrent);
                mCurrent = mCurrent.left;
            }

            TreeNode mNode = mStack.pop();

            mAnswer.add(mNode.val);
            mCurrent = mNode.right;
        }

        return mAnswer;
    }
}
