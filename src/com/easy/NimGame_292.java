/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class NimGame_292 {
    // 2019年7月15日
    // Judge mod 4's remainder equal 0 or not
    // since competitor can choose 1->3, 2->2 , 3->1 to keep always win
    // Time Complexity: O(1)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 32.8 MB, less than 5.11%
    public boolean canWinNim(int n) {
        return (n % 4) > 0;
    }
}
