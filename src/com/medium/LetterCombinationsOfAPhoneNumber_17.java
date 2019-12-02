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
public class LetterCombinationsOfAPhoneNumber_17 {
    // 2019年8月7日
    // Use DFS
    // Time Complexity: O(4^n)
    // Space Complexity:O(4^n + n)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 36.2 MB, less than 98.63%
    private static final String[] KEYS = { "", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz" };

    public List<String> letterCombinations(String digits) {
        List<String> mResult = new ArrayList<String>();
        if (digits == null || digits.length() == 0) {
            return mResult;
        }
        StringBuilder mBuilder = new StringBuilder();
        DFS(digits, 0, mBuilder, mResult);

        return mResult;
    }

    private void DFS(String digits, int mOffset, StringBuilder mBuilder, List<String> mResult) {
        if (mOffset == digits.length()) {
            mResult.add(mBuilder.toString());
            return;
        }

        String mLetters = KEYS[digits.charAt(mOffset) - '0'];
        for (int i = 0; i < mLetters.length(); i++) {
            int mLength = mBuilder.length();
            DFS(digits, mOffset + 1, mBuilder.append(mLetters.charAt(i)), mResult);
            mBuilder.setLength(mLength);
        }
    }

    // 2019年12月2日
    // Use DFS 2nd
    // Time Complexity: O(4^n)
    // Space Complexity:O(4^n + n)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 35.9 MB, less than 98.63%
    public List<String> letterCombinations4(String digits) {
        List<String> mResult = new ArrayList<>();
        if(digits == null || digits.length() == 0) {
            return mResult;
        }
        DFS(digits, 0, "", mResult);
        
        return mResult;
    }
    
    private void DFS(String digits,int mOffset, String tmp, List<String> mResult) {
        if(mOffset == digits.length()) {
            mResult.add(tmp);
            return;
        }
        String mLetters = KEYS[digits.charAt(mOffset) - '0'];
        for(char c: mLetters.toCharArray()) {
            DFS(digits, mOffset + 1, tmp + c, mResult);
        }
    }

    // 2019年8月7日
    // Use BFS
    // Time Complexity: O(4^n)
    // Space Complexity:O(2*(4^n))
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 36.2 MB, less than 98.63%
    public List<String> letterCombinations2(String digits) {
        LinkedList<String> mResult = new LinkedList<>();
        if (digits.isEmpty()) {
            return mResult;
        }
        mResult.add(""); // avoid null pointer
        while (mResult.peek().length() != digits.length()) {
            String mRemove = mResult.remove();
            String mMap = KEYS[digits.charAt(mRemove.length()) - '0'];
            for (char c : mMap.toCharArray()) {
                mResult.addLast(mRemove + c);
            }
        }
        return mResult;
    }

    // 2019年12月2日
    // Use BFS 2nd
    // Time Complexity: O(4^n)
    // Space Complexity:O(2*(4^n))
    // Runtime: 1 ms, faster than 59.45%
    // Memory Usage: 35.7 MB, less than 98.63%
    public List<String> letterCombinations3(String digits) {
        LinkedList<String> mResult = new LinkedList<>();
        if (digits.isEmpty()) {
            return mResult;
        }
        mResult.add(""); // avoid null pointer

        for (char digit : digits.toCharArray()) {
            LinkedList<String> mTmp = new LinkedList<>();
            for (String mStr : mResult) {
                for (char c : KEYS[digit - '0'].toCharArray()) {
                    mTmp.add(mStr + c);
                }
            }
            mResult = (LinkedList<String>) mTmp.clone();
            mTmp.clear();
        }
        return mResult;
    }
}
