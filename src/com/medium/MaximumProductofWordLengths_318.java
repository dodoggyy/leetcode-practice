/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class MaximumProductofWordLengths_318 {

    /**
     * @param args
     */
    public static void main(String[] args) {
        // TODO Auto-generated method stub

    }

    // Runtime: 4 ms, faster than 100.00%
    // Memory Usage: 38.4 MB, less than 93.48%
    public int maxProduct(String[] words) {
        int mResult = 0;
        int mWordLength = words.length;
        int[] mValue = new int[mWordLength];
        for (int i = 0; i < mWordLength; i++) {
            for (char c : words[i].toCharArray()) {
                mValue[i] |= 1 << (c - 'a');
            }
        }
        for (int i = 0; i < mWordLength; i++) {
            for (int j = i + 1; j < mWordLength; j++) {
                if ((mValue[i] & mValue[j]) == 0) { // two words do not share common letters
                    mResult = Math.max(mResult, words[i].length() * words[j].length());
                }
            }
        }
        return mResult;
    }
}
