/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class TrimABinarySearchTree_669 {
    // Recursive version
    // Time Complexity: O(n)
    // Space Complexity: O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 38.1 MB, less than 95.57%
    public TreeNode trimBST(TreeNode root, int L, int R) {
        if (root == null) {
            return null;
        }
        // case 1: root < L
        if (root.val < L) {
            return trimBST(root.right, L, R);
        }
        // case 2: root > R
        if (root.val > R) {
            return trimBST(root.left, L, R);
        }
        // case 3: L <= root <= R
        root.left = trimBST(root.left, L, R);
        root.right = trimBST(root.right, L, R);

        return root;
    }
}
