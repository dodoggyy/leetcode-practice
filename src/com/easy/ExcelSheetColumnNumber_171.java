/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ExcelSheetColumnNumber_171 {
    public static void main(String[] args) {
        String s = "Z";
        ExcelSheetColumnNumber_171 mTest = new ExcelSheetColumnNumber_171();
        System.out.println(mTest.titleToNumber(s));
    }

    // 2019年7月10日
    // Use digit judgment
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 35.9 MB, less than 99.68%
    public int titleToNumber(String s) {
        int mResult = 0;
        for (int i = 0; i < s.length(); i++) {
            mResult += (s.charAt(i) - 'A' + 1) * Math.pow(26, (s.length() - i - 1));
        }

        return mResult;
    }
}
