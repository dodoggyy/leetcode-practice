/**
 * 
 */
package com.easy;

import java.util.Arrays;
import java.util.Stack;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ShortestUnsortedContinuousSubarray_581 {
    // 2019年7月26日
    // Use sorting
    // Time Complexity: O(nlogn)
    // Space Complexity:O(n)
    // Runtime: 7 ms, faster than 56.59%
    // Memory Usage: 40.5 MB, less than 95.60%
    public int findUnsortedSubarray(int[] nums) {
        int[] mTmp = nums.clone();
        Arrays.sort(mTmp);

        boolean bIsSame = true;
        int mIndexLeft = 0, mIndexRight = nums.length - 1;
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] != mTmp[i]) {
                mIndexLeft = i;
                bIsSame = false;
                break;
            }
        }
        for (int i = nums.length - 1; i >= 0; i--) {
            if (nums[i] != mTmp[i]) {
                mIndexRight = i;
                bIsSame = false;
                break;
            }
        }

        return bIsSame ? 0 : mIndexRight - mIndexLeft + 1;
    }

    // 2019年7月26日
    // Use stack
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 50 ms, faster than 12.72%
    // Memory Usage: 40.8 MB, less than 84.35%
    public int findUnsortedSubarray2(int[] nums) {
        Stack<Integer> mStack = new Stack<>();
        int mLeft = nums.length, mRight = 0;
        for (int i = 0; i < nums.length; i++) {
            while (!mStack.isEmpty() && nums[mStack.peek()] > nums[i]) {
                mLeft = Math.min(mLeft, mStack.pop());
            }
            mStack.push(i);
        }
        mStack.clear();
        for (int i = nums.length - 1; i >= 0; i--) {
            while (!mStack.isEmpty() && nums[mStack.peek()] < nums[i]) {
                mRight = Math.max(mRight, mStack.pop());
            }
            mStack.push(i);
        }

        return mRight - mLeft > 0 ? mRight - mLeft + 1 : 0;
    }

    // 2019年7月26日
    // Without Using Extra Space
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 2 ms, faster than 92.21%
    // Memory Usage: 39.6 MB, less than 98.53%
    public int findUnsortedSubarray3(int[] nums) {
        int mLength = nums.length;
        int mLeft = nums.length, mRight = 0;
        int mN = nums[mLength - 1], mX = nums[0];
        for (int i = 1; i < mLength; i++) {
            mX = Math.max(mX, nums[i]);
            mN = Math.min(mN, nums[mLength - 1 - i]);
            if (mX > nums[i]) {
                mRight = i;
            }
            if (mN < nums[mLength - 1 - i]) {
                mLeft = mLength - 1 - i;
            }
        }
        return mRight - mLeft > 0 ? mRight - mLeft + 1 : 0;
    }
}
