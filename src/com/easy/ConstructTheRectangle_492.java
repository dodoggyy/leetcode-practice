/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ConstructTheRectangle_492 {
    // 2019年7月23日
    // Use Math factor judgement
    // Time Complexity: O(n^1/2)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 33.6 MB, less than 5.34%
    public int[] constructRectangle(int area) {
        for (int i = (int) Math.sqrt(area); i > 0; i--) {
            if (area % i == 0) {
                return new int[] { Math.max(area / i, i), Math.min(area / i, i) };
            }
        }

        return null;
    }
}
