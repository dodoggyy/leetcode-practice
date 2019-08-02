/**
 * 
 */
package com.easy;


/**
 * @author Chris Lin
 * @version 1.0
 */
public class SubtreeOfAnotherTree_572 {
    
    // Recursive version
    // Time Complexity O(N^2), m for numbers of tree and n for numbers of subtree
    // Runtime: 5 ms, faster than 98.89%
    // Memory Usage: 41.3 MB, less than 88.99%
    public boolean isSubtree(TreeNode s, TreeNode t) {
        if(s == null) {
            return false;
        }
        return isSubTreeWithRoot(s, t) || isSubtree(s.left, t) || isSubtree(s.right, t);
    }

    private boolean isSubTreeWithRoot(TreeNode s, TreeNode t) {
        if (s == null && t == null) {
            return true;
        }
        if (s == null || t == null) {
            return false;
        }
        if (s.val != t.val) {
            return false;
        }
        return isSubTreeWithRoot(s.left, t.left) && isSubTreeWithRoot(s.right, t.right);
    }
    
    // use inorder traversal and judge String for substring or not
    // Time Complexity O(N)
    // Runtime: 10 ms, faster than 64.02%
    // Memory Usage: 45.1 MB, less than 5.41%
    public boolean isSubtree2(TreeNode s, TreeNode t) {
        String tree = preorder(s);
        String subTree = preorder(t);
        return (tree.indexOf(subTree) >= 0);
    }
    
    private String preorder(TreeNode mNode) {
        if(mNode == null) {
            return "#";
        }
        return "*" + mNode.val + " " + preorder(mNode.left) + " " + preorder(mNode.right);
    }
}
