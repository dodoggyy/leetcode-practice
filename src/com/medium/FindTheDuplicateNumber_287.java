/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class FindTheDuplicateNumber_287 {
    // 2019年7月27日
    // Use fast slow pointer to find cycle
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 0.0 MB, less than 100%
    public int findDuplicate(int[] nums) {
        int mFast = 0;
        int mSlow = 0;

        // find cycle
        do {
            mFast = nums[nums[mFast]]; // equal mFast = mFast.next.next
            mSlow = nums[mSlow]; // equal mSlow = mSlow.next
        } while (mFast != mSlow);

        mFast = 0;
        // find cycle entry
        while (mFast != mSlow) {
            mFast = nums[mFast];
            mSlow = nums[mSlow];
        }
        return mFast;
    }
}
