/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ConvertSortedArraytoBinarySearchTree_108 {
    // Use binary search concept
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 37.7 MB, less than 54.07%
    public TreeNode sortedArrayToBST(int[] nums) {
        return toBST(nums, 0, nums.length - 1);
    }

    private TreeNode toBST(int[] nums, int mIndexLow, int mIndexHigh) {
        if (mIndexLow > mIndexHigh) {
            return null;
        }
        int mIndexMid = (mIndexLow + mIndexHigh) / 2;
        TreeNode root = new TreeNode(nums[mIndexMid]);
        root.left = toBST(nums, mIndexLow, mIndexMid - 1);
        root.right = toBST(nums, mIndexMid + 1, mIndexHigh);

        return root;
    }
}
