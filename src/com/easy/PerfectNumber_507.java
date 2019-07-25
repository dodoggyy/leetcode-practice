/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class PerfectNumber_507 {
    public static void main(String[] args) {
        PerfectNumber_507 mTest = new PerfectNumber_507();
        int num = 6;
        System.out.println(mTest.checkPerfectNumber(num));
    }

    // 2019年7月25日
    // Use iterative
    // Time Complexity: O(n^1/2)
    // Space Complexity:O(1)
    // Runtime: 2 ms, faster than 62.94%
    // Memory Usage: 33.2 MB, less than 5.04%
    public boolean checkPerfectNumber(int num) {
        if (num <= 1) {
            return false;
        }
        int mSum = 0;
        for (int i = 1; i <= (int) Math.sqrt(num); i++) {
            if (num % i == 0) {
                mSum += i;
                if (i != 1) {
                    mSum += num / i;
                }
            }
        }
        return mSum == num;
    }
}
