/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class SerializeAndDeserializeBST_449 {
    // 2019年7月30日
    // Use DFS
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 5 ms, faster than 92.70%
    // Memory Usage: 37.7 MB, less than 99.79%
    public class Codec {

        // Encodes a tree to a single string.
        public String serialize(TreeNode root) {
            StringBuilder mBuilder = new StringBuilder();
            preorder(root, mBuilder);
            return mBuilder.toString();
        }

        private void preorder(TreeNode root, StringBuilder mBuilder) {
            if (root == null) {
                // mBuilder.append("#,");
            } else {
                mBuilder.append(root.val + ",");
                preorder(root.left, mBuilder);
                preorder(root.right, mBuilder);
            }
        }

        // Decodes your encoded data to tree.
        public TreeNode deserialize(String data) {
            TreeNode root = null;
            for (String mStr : data.split(",")) {
                if (mStr.length() > 0) {
                    root = buildBST(root, Integer.parseInt(mStr));
                }
            }
            return root;
        }

        private TreeNode buildBST(TreeNode root, int mValue) {
            if (root == null) {
                return new TreeNode(mValue);
            }
            if (mValue < root.val) {
                root.left = buildBST(root.left, mValue);
            } else {
                root.right = buildBST(root.right, mValue);
            }
            return root;
        }
    }

}
