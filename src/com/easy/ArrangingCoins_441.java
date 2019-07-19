/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ArrangingCoins_441 {
    // 2019年7月20日
    // Use iterative minus
    // Time Complexity: O(n^1/2)
    // Space Complexity:O(1)
    // Runtime: 8 ms, faster than 33.22%
    // Memory Usage: 33.6 MB, less than 5.39%
    public int arrangeCoins(int n) {
        int mRow = 1;
        while (n > 0) {
            n -= ++mRow;
        }

        return mRow - 1;
    }

    // 2019年7月20日
    // Use Math formula
    // Time Complexity: O(1)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 33.5 MB, less than 5.39%
    public int arrangeCoins2(int n) {
        // x*(x+1)/2 <= n
        // => x^2 + x <= 2n
        // => 4*x^2 + 4x + 1 <= 4*2n +1
        // => (2x + 1)^2 <= 8n + 1
        // => 2x + 1 <= sqrt(8n + 1)
        // => x <= (sqrt(8n + 1)- 1)/2

        return (int) (Math.sqrt(8 * (long) n + 1) - 1) / 2;
    }
}
