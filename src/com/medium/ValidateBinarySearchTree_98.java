package com.medium;

import java.util.ArrayList;
import java.util.List;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ValidateBinarySearchTree_98 {
    // 2019年9月12日
    // Use Recursive
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 1 ms, faster than 48.17%
    // Memory Usage: 39.3 MB, less than 80.47%
    public boolean isValidBST(TreeNode root) {
        return checkBST(root, null, null);
    }

    private boolean checkBST(TreeNode root, Integer mMin, Integer mMax) {
        if (root == null) {
            return true;
        }
        int mValue = root.val;
        if (mMin != null && mValue <= mMin) {
            return false;
        }
        if (mMax != null && mValue >= mMax) {
            return false;
        }
        if (!checkBST(root.left, mMin, mValue)) {
            return false;
        }
        if (!checkBST(root.right, mValue, mMax)) {
            return false;
        }

        return true;
    }

    // 2019年9月12日
    // Use inorder traversal
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 2 ms, faster than 24.31%
    // Memory Usage: 37.2 MB, less than 100.00%
    private List<Integer> mList;

    public boolean isValidBST2(TreeNode root) {
        mList = new ArrayList<>();
        copyBST(root);
        for (int i = 0; i < mList.size() - 1; i++) {
            if (mList.get(i) >= mList.get(i + 1)) {
                return false;
            }
        }
        return true;
    }

    private void copyBST(TreeNode root) {
        if (root == null) {
            return;
        }
        copyBST(root.left);
        mList.add(root.val);
        copyBST(root.right);
    }
}
