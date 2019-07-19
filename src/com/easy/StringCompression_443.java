/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class StringCompression_443 {
    // 2019年7月20日
    // Use iterative parsing
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 96.95%
    // Memory Usage: 38 MB, less than 5.05%
    public int compress(char[] chars) {
        int mLength = chars.length;
        int mResult = 0;
        for (int i = 1; i <= mLength; i++) {
            int mCount = 1;
            while (i < mLength && chars[i] == chars[i - 1]) {
                i++;
                mCount++;
            }
            chars[mResult++] = chars[i - 1];
            if (mCount == 1) {
                continue;
            }
            for (char c : String.valueOf(mCount + "").toCharArray()) {
                chars[mResult++] = c;
            }

        }

        return mResult;
    }
}
