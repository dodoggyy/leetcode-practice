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
public class KthSmallestElementInABST_230 {
    // Recursive version with memorize
    // Time Complexity: O(n)
    // Space Complexity: O(n)
    // Runtime: 3 ms, faster than 10.50%
    // Memory Usage: 38.3 MB, less than 95.53%
    private List<Integer> mList = new ArrayList<>();

    public int kthSmallest(TreeNode root, int k) {
        inorder(root);
        return mList.get(k);
    }

    // inorder traversal: LDR
    private void inorder(TreeNode root) {
        if (root == null) {
            return;
        }
        inorder(root.left);
        mList.add(root.val);
        inorder(root.right);
    }

    // Recursive version without memorize
    // Time Complexity: O(n)
    // Space Complexity: O(n)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 37.6 MB, less than 99.71%
    private int mAnswer = 0;
    private int mCnt = 0;

    public int kthSmallest2(TreeNode root, int k) {
        inorder(root, k);

        return mAnswer;
    }

    private void inorder(TreeNode root, int k) {
        if (root == null) {
            return;
        }
        inorder(root.left, k);
        mCnt++;
        if (k == mCnt) {
            mAnswer = root.val;
            return;
        }
        inorder(root.right, k);
    }
}
