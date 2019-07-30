/**
 * 
 */
package com.hard;

import java.util.LinkedList;
import java.util.Queue;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class SerializeAndDeserializeBinaryTree_297 {
    // 2019年7月30日
    // Use DFS(preorder traversal)
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 12 ms, faster than 60.34%
    // Memory Usage: 39.9 MB, less than 40.96%
    public class Codec {

        // Encodes a tree to a single string.
        public String serialize(TreeNode root) {
            if (root == null) {
                return "#,";
            }
            return root.val + "," + serialize(root.left) + serialize(root.right);
        }

        // Decodes your encoded data to tree.
        public TreeNode deserialize(String data) {
            Queue<TreeNode> mQueue = new LinkedList<>();
            for (String mStr : data.split(",")) {
                if (mStr.equals("#")) {
                    mQueue.offer(null);
                } else {
                    mQueue.offer(new TreeNode(Integer.valueOf(mStr)));
                }
            }
            return preorder(mQueue);
        }

        int mIndex = -1;

        private TreeNode preorder(Queue<TreeNode> mQueue) {
            if (mQueue.isEmpty()) {
                return null;
            }
            TreeNode mNode = mQueue.poll();
            if (mNode == null) {
                return null;
            }
            mNode.left = preorder(mQueue);
            mNode.right = preorder(mQueue);

            return mNode;
        }
    }

    // 2019年7月30日
    // Use BFS
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 12 ms, faster than 60.34%
    // Memory Usage: 39.9 MB, less than 40.96%
    public class Codec2 {

        // Encodes a tree to a single string.
        public String serialize(TreeNode root) {
            if (root == null) {
                return "";
            }
            StringBuilder mResult = new StringBuilder();
            Queue<TreeNode> mQueue = new LinkedList<>();
            mQueue.offer(root);

            while (!mQueue.isEmpty()) {
                TreeNode mNode = mQueue.poll();
                if (mNode == null) {
                    mResult.append("#,");
                } else {
                    mResult.append(mNode.val + ",");
                    mQueue.add(mNode.left);
                    mQueue.add(mNode.right);
                }

            }

            return mResult.toString();
        }

        // Decodes your encoded data to tree.
        public TreeNode deserialize(String data) {
            if (data.length() == 0) {
                return null;
            }
            Queue<TreeNode> mQueue = new LinkedList<>();
            String[] mValues = data.split(",");
            TreeNode root = new TreeNode(Integer.parseInt(mValues[0]));
            mQueue.add(root);
            for (int i = 1; i < mValues.length; i++) {
                TreeNode mNode = mQueue.poll();
                if (!mValues[i].equals("#")) {
                    TreeNode mLeft = new TreeNode(Integer.parseInt(mValues[i]));
                    mNode.left = mLeft;
                    mQueue.add(mLeft);
                }
                i++;
                if (!mValues[i].equals("#")) {
                    TreeNode mRight = new TreeNode(Integer.parseInt(mValues[i]));
                    mNode.right = mRight;
                    mQueue.add(mRight);
                }
            }
            return root;
        }

    }
}
