/**
 * 
 */
package com.hard;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Stack;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class BinaryTreePostorderTraversal_145 {
    // Recursive version (LRD)
    // Time Complexity: O(n)
    // Space Complexity: O(n)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 34.7 MB, less than 100.00%
    private List<Integer> mAnswer = new ArrayList<>();

    public List<Integer> postorderTraversal(TreeNode root) {
        if (root == null) {
            return mAnswer;
        }
        postorderTraversal(root.left);
        postorderTraversal(root.right);
        mAnswer.add(root.val);

        return mAnswer;
    }

    // Iterative version (LRD)
    // Time Complexity: O(n)
    // Space Complexity: O(n)
    // Runtime: 1 ms, faster than 61.18%
    // Memory Usage: 34.9 MB, less than 100.00%
    public List<Integer> postorderTraversal2(TreeNode root) {
        Stack<TreeNode> mStack = new Stack<>();
        mStack.push(root);
        while (!mStack.isEmpty()) {
            TreeNode mNode = mStack.pop();
            if (mNode == null) {
                continue;
            }
            mAnswer.add(mNode.val);
            mStack.push(mNode.left);
            mStack.push(mNode.right);
        }
        Collections.reverse(mAnswer);

        return mAnswer;
    }
}
