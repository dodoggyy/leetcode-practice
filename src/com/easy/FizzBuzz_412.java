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
public class FizzBuzz_412 {
    // 2019年7月18日
    // Use iterative
    // Time Complexity: O(n)
    // Space Complexity:O()
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 37.4 MB, less than 99.94%
    public List<String> fizzBuzz(int n) {
        List<String> mResult = new ArrayList<>();
        if (n < 1) {
            return mResult;
        }
        for (int i = 1; i <= n; i++) {
            if ((i % 3 == 0) && (i % 5 == 0)) {
                mResult.add("FizzBuzz");
            } else if (i % 3 == 0) {
                mResult.add("Fizz");
            } else if (i % 5 == 0) {
                mResult.add("Buzz");
            } else {
                mResult.add(i + "");
            }
        }
        return mResult;
    }
}
