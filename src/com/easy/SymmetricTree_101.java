/**
 * 
 */
package com.easy;

import java.util.LinkedList;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class SymmetricTree_101 {

    // Recursive version
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 37 MB, less than 92.16%
    public boolean isSymmetric(TreeNode root) {
        if (root == null) {
            return true;
        }

        return isSymmetric(root.left, root.right);
    }

    private boolean isSymmetric(TreeNode t1, TreeNode t2) {
        if (t1 == null && t2 == null) {
            return true;
        }
        if (t1 == null || t2 == null) {
            return false;
        }
        if (t1.val != t2.val) {
            return false;
        }

        return isSymmetric(t1.left, t2.right) && isSymmetric(t1.right, t2.left);
    }

    // Iterative version
    // Runtime: 1 ms, faster than 72.80%
    // Memory Usage: 40 MB, less than 8.34%
    public boolean isSymmetric2(TreeNode root) {
        if (root == null) {
            return true;
        }
        if (root.left == null && root.right == null) {
            return true;
        }
        if (root.left == null || root.right == null) {
            return false;
        }
        LinkedList<TreeNode> mQueue1 = new LinkedList<TreeNode>();
        LinkedList<TreeNode> mQueue2 = new LinkedList<TreeNode>();

        mQueue1.add(root.left);
        mQueue2.add(root.right);

        while (!mQueue1.isEmpty() && !mQueue2.isEmpty()) {
            TreeNode t1 = mQueue1.poll();
            TreeNode t2 = mQueue2.poll();

            if (t1.val != t2.val) {
                return false;
            }
            if ((t1.left != null && t2.right == null) || t1.right != null && t2.left == null) {
                return false;
            }
            if ((t1.left == null && t2.right != null) || t1.right == null && t2.left != null) {
                return false;
            }

            if (t1.left != null && t2.right != null) {
                mQueue1.push(t1.left);
                mQueue2.push(t2.right);
            }
            if (t1.right != null && t2.left != null) {
                mQueue1.push(t1.right);
                mQueue2.push(t2.left);
            }
        }
        return true;
    }
}
