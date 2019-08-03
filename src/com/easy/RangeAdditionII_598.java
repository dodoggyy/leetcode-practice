/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class RangeAdditionII_598 {
    // 2019年8月3日
    // Use simple minimum judgment
    // Time Complexity: O(x)
    // Space Complexity:O(1)
    // x: length of operation
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 36.1 MB, less than 96.55%
    public int maxCount(int m, int n, int[][] ops) {
        for (int[] op : ops) {
            m = Math.min(m, op[0]);
            n = Math.min(n, op[1]);
        }
        return m * n;
    }
}
