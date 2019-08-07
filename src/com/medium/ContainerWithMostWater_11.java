/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ContainerWithMostWater_11 {
    // 2019年8月7日
    // Use 2 pointers
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 2 ms, faster than 94.39%
    // Memory Usage: 39.9 MB, less than 93.11%
    public int maxArea(int[] height) {
        int mIndexLeft = 0, mIndexRight = height.length - 1;
        int mResult = 0;
        while(mIndexLeft < mIndexRight) {
            int mHeight = Math.min(height[mIndexLeft], height[mIndexRight]);
            mResult = Math.max(mResult, mHeight* (mIndexRight - mIndexLeft));
            if(height[mIndexLeft] < height[mIndexRight]) {
                mIndexLeft++;
            } else {
                mIndexRight--;
            }
        }
        return mResult;
    }
}
