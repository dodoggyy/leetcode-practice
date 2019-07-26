/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ProductOfArrayExceptSelf_238 {
    public static void main(String[] args) {
        ProductOfArrayExceptSelf_238 mTest = new ProductOfArrayExceptSelf_238();
        int[] nums = new int[] { 4, 5, 1, 8, 2 };
        mTest.productExceptSelf2(nums);
    }

    // 2019年7月26日
    // Use brute force
    // Time Complexity: O(n^2)
    // Space Complexity:O(1)
    // Time Limit Exceeded
    public int[] productExceptSelf(int[] nums) {
        if (nums == null || nums.length == 0) {
            return new int[] {};
        }
        int mLength = nums.length;
        int[] mResult = new int[mLength];
        for (int i = 0; i < mLength; i++) {
            int mTmp = 1;
            for (int j = 0; j < mLength; j++) {
                if (j != i) {
                    mTmp *= nums[j];
                }
            }
            mResult[i] = mTmp;
        }

        return mResult;
    }

    // 2019年7月26日
    // Use dual array record
    // Time Complexity: O(3n) = O(n)
    // Space Complexity:O(2n) = O(n)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 41 MB, less than 99.95%
    public int[] productExceptSelf2(int[] nums) {
        if (nums == null || nums.length == 0) {
            return new int[] {};
        }
        int mLength = nums.length;
        int[] mResult = new int[mLength];
        int[] mLeft = new int[mLength];
        int[] mRight = new int[mLength];

        mLeft[0] = 1;
        mRight[mLength - 1] = 1;
        for (int i = 1; i < mLength; i++) {
            mLeft[i] = mLeft[i - 1] * nums[i - 1];
        }

        for (int i = mLength - 2; i >= 0; i--) {
            mRight[i] = mRight[i + 1] * nums[i + 1];
        }

        for (int i = 0; i < mLength; i++) {
            mResult[i] = mLeft[i] * mRight[i];
        }

        // for (int i = 0; i < mLength; i++) {
        // System.out.print(mLeft[i] + " ");
        // }
        // System.out.println();
        // for (int i = 0; i < mLength; i++) {
        // System.out.print(mRight[i] + " ");
        // }

        return mResult;
    }

    // 2019年7月26日
    // Use 1 parameter and result array record
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 41.7 MB, less than 99.95%
    public int[] productExceptSelf3(int[] nums) {
        if (nums == null || nums.length == 0) {
            return new int[] {};
        }
        int mLength = nums.length;
        int[] mResult = new int[mLength];

        mResult[0] = 1;

        for (int i = 1; i < mLength; i++) {
            mResult[i] = mResult[i - 1] * nums[i - 1];
        }

        int R = 1;
        for (int i = mLength - 1; i >= 0; i--) {
            mResult[i] = mResult[i] * R;
            R *= nums[i];
        }

        return mResult;
    }
}
