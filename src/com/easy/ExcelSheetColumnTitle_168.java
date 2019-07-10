/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ExcelSheetColumnTitle_168 {
    public static void main(String[] args) {
        ExcelSheetColumnTitle_168 mTest = new ExcelSheetColumnTitle_168();
        System.out.println(mTest.convertToTitle(1));
    }

    // 2019年7月10日
    // Use digit judgment
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 33.9 MB, less than 100.00%
    public String convertToTitle(int n) {
        StringBuilder mResult = new StringBuilder();
        while (n > 0) {
            int mRemain = n % 26;

            // System.out.println(n);

            if (mRemain == 0) {
                mResult.insert(0, "Z");
                n = (n / 26) - 1;
            } else {
                mResult.insert(0, (char) (mRemain + 'A' - 1));
                n /= 26;
            }

        }
        return mResult.toString();
    }

}
