/**
 * 
 */
package com.easy;

import java.util.HashSet;
import java.util.Set;
import java.util.Stack;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ConstructStringFromBinaryTree_606 {
    // 2019年8月4日
    // Use recursive
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 9 ms, faster than 18.49%
    // Memory Usage: 44.4 MB, less than 6.31%
    public String tree2str(TreeNode t) {
        if (t == null) {
            return "";
        }
        String mData = t.val + "";
        String mLeft = tree2str(t.left);
        String mRight = tree2str(t.right);
        if (t.left == null && t.right == null) {
            return mData;
        }

        if (t.right == null) {
            return mData + "(" + mLeft + ")";
        }
        return mData + "(" + mLeft + ")" + "(" + mRight + ")";
    }

    // 2019年8月4日
    // Use iterative via stack
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 7 ms, faster than 40.14%
    // Memory Usage: 38.3 MB, less than 98.20%
    public String tree2str2(TreeNode t) {
        if (t == null) {
            return "";
        }
        Stack<TreeNode> mStack = new Stack<>();
        Set<TreeNode> mSet = new HashSet<>();
        StringBuilder mResult = new StringBuilder();
        mStack.push(t);
        while (!mStack.isEmpty()) {
            TreeNode mNode = mStack.peek();
            if (mSet.contains(mNode)) {
                mStack.pop();
                mResult.append(")");
            } else {
                mSet.add(mNode);
                mResult.append("(" + mNode.val);
                if (mNode.left == null && mNode.right != null) {
                    mResult.append("()");
                }
                if (mNode.right != null) {
                    mStack.push(mNode.right);
                }
                if (mNode.left != null) {
                    mStack.push(mNode.left);
                }
            }
        }
        return mResult.substring(1, mResult.length() - 1);
    }
}
