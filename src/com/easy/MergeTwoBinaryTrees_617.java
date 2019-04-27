/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class MergeTwoBinaryTrees_617 {
    
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 41.4 MB, less than 55.07%
    public TreeNode mergeTrees(TreeNode t1, TreeNode t2) {
        if(t1 == null && t2 == null) {
            return null;
        }
        if(t1 == null) {
            return t2;
        }
        if(t2 == null) {
            return t1;
        }
        TreeNode mMergeTree = new TreeNode(t1.val + t2.val);
        mMergeTree.left = mergeTrees(t1.left, t2.left);
        mMergeTree.right = mergeTrees(t1.right, t2.right);
        
        return mMergeTree;
    }
}
