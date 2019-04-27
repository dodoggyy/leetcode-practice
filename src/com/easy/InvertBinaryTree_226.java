/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class InvertBinaryTree_226 {
    
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 35.7 MB, less than 29.71%
    public TreeNode invertTree(TreeNode root) {
        if(root == null) {
            return null;
        }
        TreeNode mLeftTree = root.left; // tmp left tree
        root.left = invertTree(root.right);
        root.right = invertTree(mLeftTree);
        
        return root;
    }
}
