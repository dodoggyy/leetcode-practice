/**
 * 
 */
package com.easy;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class FindModeInBinarySearchTree_501 {
    // 2019年7月24日
    // Use HashMap
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 8 ms, faster than 23.74%
    // Memory Usage: 37.8 MB, less than 99.26%
    private Map<Integer, Integer> mMap = new HashMap<>();
    private int mMax = 0;
    private List<Integer> mList = new ArrayList<>();

    public int[] findMode(TreeNode root) {
        inorder(root);

        for (int key : mMap.keySet()) {
            if (mMap.get(key) == mMax) {
                mList.add(key);
            }
        }

        int[] mResult = new int[mList.size()];
        int mIndex = 0;
        for (int num : mList) {
            mResult[mIndex++] = num;
        }

        return mResult;
    }

    private void inorder(TreeNode root) {
        if (root == null) {
            return;
        }
        inorder(root.left);
        visit(root.val);
        inorder(root.right);
    }

    private void visit(int mValue) {
        mMap.put(mValue, mMap.getOrDefault(mValue, 0) + 1);
        mMax = Math.max(mMap.get(mValue), mMax);
    }

    // 2019年7月24日
    // Use count via inorder traversal
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 97.26%
    // Memory Usage: 37.6 MB, less than 100.00%
    private int mCount = 0;
    private int mMaxCount = -1;
    private TreeNode mPre;
    private List<Integer> mModes = new ArrayList<>();

    public int[] findMode2(TreeNode root) {
        if (root == null) {
            return new int[0];
        }
        mPre = root;
        inorder2(root);

        int[] mResult = new int[mModes.size()];
        for (int i = 0; i < mModes.size(); i++) {
            mResult[i] = mModes.get(i);
        }

        return mResult;
    }

    private void inorder2(TreeNode root) {
        if (root == null) {
            return;
        }
        inorder2(root.left);
        mCount = (mPre.val == root.val) ? mCount + 1 : 1;

        if (mCount == mMaxCount) {
            mModes.add(root.val);
        } else if (mCount > mMaxCount) {
            mModes.clear();
            mModes.add(root.val);
            mMaxCount = mCount;
        }

        mPre = root;
        inorder2(root.right);
    }
}
