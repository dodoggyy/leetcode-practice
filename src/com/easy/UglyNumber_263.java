/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class UglyNumber_263 {
    // 2019年7月15日
    // Simple solution
    // Time Complexity: O(logn)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 33.6 MB, less than 5.08%
    public boolean isUgly(int num) {
        if (num == 0) {
            return false;
        }

        while (num > 1) {
            int mTmp = num;
            if (num % 2 == 0) {
                num /= 2;
            }
            if (num % 3 == 0) {
                num /= 3;
            }
            if (num % 5 == 0) {
                num /= 5;
            }
            if (mTmp == num) {
                break;
            }
        }

        return num == 1;
    }

    // 2019年7月15日
    // Refine simple solution
    // Time Complexity: O(logn)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 33.3 MB, less than 5.08%
    public boolean isUgly2(int num) {
        if (num == 0) {
            return false;
        }
        while (num != 1) {
            if (num % 2 == 0) {
                num /= 2;
            } else if (num % 3 == 0) {
                num /= 3;
            } else if (num % 5 == 0) {
                num /= 5;
            } else {
                return false;
            }
        }

        return true;
    }
}
