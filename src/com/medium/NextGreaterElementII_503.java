/**
 * 
 */
package com.medium;

import java.util.Arrays;
import java.util.Stack;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class NextGreaterElementII_503 {
    // Use stack with order
    // Time Complexity: O(2n) = O(n)
    // Space Complexity:O(n)
    // Runtime: 24 ms, faster than 83.32%
    // Memory Usage: 41.3 MB, less than 99.91%
    public int[] nextGreaterElements(int[] nums) {
        int mLength = nums.length;
        int[] mResult = new int[mLength];
        Arrays.fill(mResult, -1);
        Stack<Integer> mStack = new Stack<>();

        for (int i = 0; i < 2 * mLength; i++) {
            int mNum = nums[i % mLength];
            while (!mStack.isEmpty() && nums[mStack.peek()] < mNum) {
                mResult[mStack.pop()] = mNum;
            }
            if (i < mLength) {
                mStack.push(i);
            }
        }

        return mResult;
    }

    // Use stack with reverse order
    // Time Complexity: O(2n) = O(n)
    // Space Complexity:O(n)
    // Runtime: 44 ms, faster than 34.15%
    // Memory Usage: 40.6 MB, less than 100.00%
    public int[] nextGreaterElements2(int[] nums) {
        if (nums == null || nums.length == 0) {
            return nums;
        }
        if (nums.length == 1) {
            return new int[] { -1 };
        }

        int mLength = nums.length;
        int[] mResult = new int[mLength];
        Arrays.fill(mResult, -1);
        Stack<Integer> mStack = new Stack<>();
        for (int i = mLength * 2 - 1; i >= 0; i--) {
            int mIndex = i % mLength;
            while (!mStack.isEmpty() && nums[mStack.peek()] <= nums[mIndex]) {
                mStack.pop();
            }
            if (!mStack.isEmpty()) {
                mResult[mIndex] = nums[mStack.peek()];
            }
            mStack.push(mIndex);
        }
        return mResult;
    }
}
