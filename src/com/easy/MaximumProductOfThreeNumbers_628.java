/**
 * 
 */
package com.easy;

import java.util.Arrays;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class MaximumProductOfThreeNumbers_628 {
    // 2020年2月3日
    // Use sorting:
    // Time Complexity: O(nlogn)
    // Space Complexity:O(1)
    // Runtime: 9 ms, faster than 65.09%
    // Memory Usage: 42.3 MB, less than 7.14%
    public int maximumProduct(int[] nums) {
        Arrays.sort(nums);
        int size = nums.length;
        return Math.max(nums[0] * nums[1] * nums[size - 1], nums[size - 1] * nums[size - 2] * nums[size - 3]);
    }

    // 2020年2月3日
    // Use Single Scan:
    // Time Complexity: O(nlogn)
    // Space Complexity:O(1)
    // Runtime: 2 ms, faster than 99.21%
    // Memory Usage: 41.9 MB, less than 7.14%
    public int maximumProduct2(int[] nums) {
        int max1 = Integer.MIN_VALUE, max2 = Integer.MIN_VALUE, max3 = Integer.MIN_VALUE;
        int min1 = Integer.MAX_VALUE, min2 = Integer.MAX_VALUE;

        for (int num : nums) {
            if (num > max1) {
                max3 = max2;
                max2 = max1;
                max1 = num;
            } else if (num > max2) {
                max3 = max2;
                max2 = num;
            } else if (num > max3) {
                max3 = num;
            }

            if (num < min1) {
                min2 = min1;
                min1 = num;
            } else if (num < min2) {
                min2 = num;
            }
        }

        return Math.max(min1 * min2 * max1, max1 * max2 * max3);
    }
}
