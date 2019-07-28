/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ReverseWordsInAStringIII_557 {
    // 2019年7月28日
    // Use build-in library
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 5 ms, faster than 68.28%
    // Memory Usage: 38.6 MB, less than 92.41%
    public String reverseWords(String s) {
        String[] mTmp = s.split(" ");
        StringBuilder mResult = new StringBuilder();
        for (String mStr : mTmp) {
            mResult.append(new StringBuilder().append(mStr).reverse().toString() + " ");
        }

        return mResult.toString().trim();
    }

    // 2019年7月28日
    // Use build-in library with other method
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 10 ms, faster than 37.05%
    // Memory Usage: 38.4 MB, less than 97.74% 
    public String reverseWords2(String s) {
        StringBuilder mResult = new StringBuilder();
        StringBuilder mTmp = new StringBuilder();
        
        for(int i = 0; i < s.length();i++) {
            if(s.charAt(i) != ' ') {
                mTmp.append(s.charAt(i));
            } else {
                mResult.append(mTmp.reverse());
                mResult.append(" ");
                mTmp.setLength(0);
            }
        }
        mResult.append(mTmp.reverse());
        
        return mResult.toString();
    }
}
