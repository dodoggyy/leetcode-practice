/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ConvertANumberToHexadecimal_405 {
    // 2019年7月18日
    // Use library method
    // Time Complexity: O(1)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 33.8 MB, less than 100.00%
    public String toHex(int num) {
        return Integer.toHexString(num);
    }

    // 2019年7月18日
    // Use bit shift comparison
    // Time Complexity: O(1ogn)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 34 MB, less than 100.00%
    public String toHex2(int num) {
        char[] mHex = { '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f' };
        StringBuilder mResult = new StringBuilder();
        if (num == 0) {
            return "0";
        }
        while (num != 0) {
            mResult.insert(0, mHex[(num&0xF)]);
//            mResult.insert(0, mHex[(num < 0 ? (num % 16 + 16) % 16 : num % 16)]);
            num >>>= 4;
        }

        return mResult.toString();
    }
}
