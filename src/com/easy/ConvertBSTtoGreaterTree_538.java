/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ConvertBSTtoGreaterTree_538 {
    // Use RDL traversal
    // Time Complexity: O(n)
    // Space Complexity: O(n)
    // Runtime: 1 ms, faster than 57.30% 
    // Memory Usage: 40.6 MB, less than 57.93%
    private int mValue = 0;
    public TreeNode convertBST(TreeNode root) {
        traversal(root);
        return root;
    }
    
    // traversal: RDL
    private void traversal(TreeNode root) {
        if(root == null) {
            return;
        }
        traversal(root.right);
        mValue += root.val;
        root.val = mValue;
        traversal(root.left);
    }
}
