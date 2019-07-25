/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class DetectCapital_520 {
    // 2019年7月26日
    // Use String comparison
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 2 ms, faster than 25.44%
    // Memory Usage: 35 MB, less than 100.00%
    public boolean detectCapitalUse(String word) {
        return word.substring(1).equals(word.substring(1).toLowerCase()) || word.equals(word.toUpperCase());
    }

    // 2019年7月26日
    // Use regular expression
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 7 ms, faster than 6.81%
    // Memory Usage: 35.9 MB, less than 100.00%
    public boolean detectCapitalUse2(String word) {
        // X? : 0 or 1
        // X* : 0 or more
        // X+ : 1 or more
        return word.matches("[A-Z]+|[A-Z]?[a-z]+");
    }
}
