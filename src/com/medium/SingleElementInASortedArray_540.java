/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class SingleElementInASortedArray_540 {
    // 2019年10月2日
    // Use XOR operation
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 37.4 MB, less than 64.00%
    public int singleNonDuplicate(int[] nums) {
        int mResult = 0;
        for (int num : nums) {
            mResult ^= num;
        }

        return mResult;
    }

    // 2019年10月2日
    // Use binary search
    // Time Complexity: O(logn)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 37.6 MB, less than 40.00%
    public int singleNonDuplicate2(int[] nums) {
        int mIndexLow = 0, mIndexHigh = nums.length - 1;

        while (mIndexLow < mIndexHigh) {
            // We want the first element of the middle pair,
            // which should be at an even index if the left part is sorted.
            // Example:
            // Index: 0 1 2 3 4 5 6
            // Array: 1 1 3 3 4 8 8
            //            ^
            int mIndexMid = mIndexLow + (mIndexHigh - mIndexLow) / 2;
            if (mIndexMid % 2 == 1) {
                mIndexMid--;
            }

            // We didn't find a pair. The single element must be on the left.
            // (pipes mean start & end)
            // Example: |0 1 1 3 3 6 6|
            //               ^ ^
            // Next:    |0 1 1|3 3 6 6
            if (nums[mIndexMid] != nums[mIndexMid + 1]) {
                mIndexHigh = mIndexMid;
            } else {
                // We found a pair. The single element must be on the right.
                // Example: |1 1 3 3 5 6 6|
                //               ^ ^
                // Next:     1 1 3 3|5 6 6|
                mIndexLow = mIndexMid + 2;
            }
        }

        // 'mIndexLow' should always be at the beginning of a pair.
        // When 'mIndexLow > mIndexHigh', start must be the single element.
        return nums[mIndexLow];
    }
}
