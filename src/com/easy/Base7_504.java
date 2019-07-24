/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class Base7_504 {
    // 2019年7月24日
    // Use build-in library
    // Time Complexity: O(1)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 93.81%
    // Memory Usage: 34.1 MB, less than 100.00%
    public String convertToBase7(int num) {
        return Integer.toString(num, 7);
    }

    // 2019年7月24日
    // Use Iterative
    // Time Complexity: O(1ogn)
    // Space Complexity:O(1)
    // Runtime: 2 ms, faster than 93.81%
    // Memory Usage: 34.4 MB, less than 100.00%
    public String convertToBase72(int num) {
        if (num == 0) {
            return 0 + "";
        }
        StringBuilder mResult = new StringBuilder();
        boolean bIsNegative = false;
        if (num < 0) {
            bIsNegative = true;
        }

        while (num != 0) {
            mResult.insert(0, Math.abs(num % 7));
            num /= 7;
        }
        if (bIsNegative) {
            mResult.insert(0, "-");
        }

        return mResult.toString();
    }

    // 2019年7月24日
    // Use Recursive
    // Time Complexity: O(1ogn)
    // Space Complexity:O(1)
    // Runtime: 2 ms, faster than 93.81%
    // Memory Usage: 34.9 MB, less than 100.00%
    public String convertToBase73(int num) {
        if (num < 0) {
            return "-" + convertToBase73(-num);
        }
        if (num < 7) {
            return Integer.toString(num);
        }
        return convertToBase73(num / 7) + Integer.toString(num % 7);
    }
}
