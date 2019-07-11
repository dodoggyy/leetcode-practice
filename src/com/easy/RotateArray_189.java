/**
 * 
 */
package com.easy;

import java.util.HashMap;
import java.util.Map;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class RotateArray_189 {
    // 2019年7月11日
    // Use extra array
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 36.4 MB, less than 99.97%
    public void rotate(int[] nums, int k) {
        int mLength = nums.length;
        int[] mTmp = nums.clone();
        for (int i = 0; i < nums.length; i++) {
            nums[(i + k) % mLength] = mTmp[i];
        }
    }

    // 2019年7月11日
    // Use HashMap
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 3 ms, faster than 25.09%
    // Memory Usage: 37.9 MB, less than 88.65%
    public void rotate2(int[] nums, int k) {
        Map<Integer, Integer> mMap = new HashMap<>();
        int mLength = nums.length;
        for (int i = 0; i < mLength; i++) {
            mMap.put(i, nums[i]);
        }
        for (int i = 0; i < mLength; i++) {
            nums[(i + k) % mLength] = mMap.get(i);
        }
    }

    // 2019年7月11日
    // Use Cyclic Replacements
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 36.3 MB, less than 99.97%
    public void rotate3(int[] nums, int k) {
        int mLength = nums.length;
        int mCount = 0;
        for (int mStart = 0; mCount < mLength; mStart++) {
            int mCurrent = mStart;
            int mPrevious = nums[mStart];
            do {
                int mNext = (mCurrent + k) % mLength;
                int mTmp = nums[mNext];
                nums[mNext] = mPrevious;
                mPrevious = mTmp;
                mCurrent = mNext;
                mCount++;
            } while (mStart != mCurrent);
        }
    }

    // 2019年7月11日
    // Use Reverse
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 37.2 MB, less than 99.97%
    public void rotate4(int[] nums, int k) {
        int mLength = nums.length;
        k %= mLength;
        reverse(nums, 0, mLength);
        reverse(nums, 0, k - 1);
        reverse(nums, k, mLength - 1);
    }

    private void reverse(int[] nums, int mStart, int mEnd) {
        while (mStart < mEnd) {
            int mTmp = nums[mStart];
            nums[mStart] = nums[mEnd];
            nums[mEnd] = mTmp;
            mStart++;
            mEnd--;
        }
    }
}
