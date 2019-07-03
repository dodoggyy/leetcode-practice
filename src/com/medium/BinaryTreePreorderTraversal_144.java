/**
 * 
 */
package com.medium;

import java.util.ArrayList;
import java.util.List;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class BinaryTreePreorderTraversal_144 {
    // Use preorder traversal (DLR)
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

}
