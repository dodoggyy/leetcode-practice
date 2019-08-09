/**
 * 
 */
package com.medium;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class GenerateParentheses_22 {
    // 2019年8月9日
    // Use DFS
    // Time Complexity: O(2^n)
    // Space Complexity:O(k + n)
    // Runtime: 1 ms, faster than 93.54%
    // Memory Usage: 36.1 MB, less than 100.00%
    public List<String> generateParenthesis(int n) {
        List<String> mResult = new ArrayList<>();

        if (n <= 0) {
            return mResult;
        }
        helper(n, 0, 0, "", mResult);

        return mResult;
    }

    private void helper(int n, int mCountLeft, int mCountRight, String mTmpStr, List<String> mResult) {
        if (mTmpStr.length() == 2 * n) {
            mResult.add(mTmpStr);
        }

        if (mCountLeft < n) {
            helper(n, mCountLeft + 1, mCountRight, mTmpStr + "(", mResult);
        }
        if (mCountRight < mCountLeft) {
            helper(n, mCountLeft, mCountRight + 1, mTmpStr + ")", mResult);
        }
    }
}
