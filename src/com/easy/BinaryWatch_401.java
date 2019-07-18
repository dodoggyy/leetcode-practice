/**
 * 
 */
package com.easy;

import java.util.ArrayList;
import java.util.List;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class BinaryWatch_401 {
    // 2019年7月18日
    // Use bit count
    // Time Complexity: O(h*m)
    // Space Complexity:O(1)
    // Runtime: 11 ms, faster than 42.52%
    // Memory Usage: 37 MB, less than 96.25%
    public List<String> readBinaryWatch(int num) {
        List<String> mResult = new ArrayList<>();

        for (int i = 0; i < 12; i++) { // hour range: 0~12
            for (int j = 0; j < 60; j++) { // minute: 0~59
                if ((Integer.bitCount(i * 64 + j) == num)) {
                    mResult.add(String.format("%d%:02d", i, j));
                }
            }
        }

        return mResult;
    }
}
