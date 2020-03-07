/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class HowManyNumbersAreSmallerThanTheCurrentNumber_1365 {
    // Use counting sort:
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 99.72%
    // Memory Usage: 42 MB, less than 100.00%
    public int[] smallerNumbersThanCurrent(int[] nums) {
        int[] result = new int[nums.length];
        int[] count = new int[101];

        for (int num : nums) {
            count[num]++;
        }

        int tmp = 0;

        for (int i = 0; i < count.length; i++) {
            int less = count[i];
            count[i] = tmp;
            tmp += less;
        }

        for (int i = 0; i < nums.length; i++) {
            result[i] = count[nums[i]];
        }

        return result;
    }
}
