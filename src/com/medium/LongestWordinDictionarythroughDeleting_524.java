/**
 * 
 */
package com.medium;

import java.util.Arrays;
import java.util.List;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class LongestWordinDictionarythroughDeleting_524 {

    public static void main(String[] args) {
        String mInput = "bab";// "abpcplea";
        List<String> mList = Arrays.asList(new String[] { "ale", "apple", "monkey", "plea" });
        List<String> mList2 = Arrays.asList(new String[] { "ba", "ab", "a", "b" });

        System.out.println("Ans: " + findLongestWord(mInput, mList2));
    }

    // Runtime: 36 ms, faster than 49.89%
    // Memory Usage: 43.1 MB, less than 7.22%
    public static String findLongestWord(String s, List<String> d) {
        if (s.length() == 0 || d.size() == 0) {
            return "";
        }

        String mFinalAnser = "";
        int mWordLength = s.length();
        for (int i = 0; i < d.size(); i++) {
            String mTmp = "";
            String mDicContent = d.get(i);
            int mDicIndex = 0, mWordIndex = 0;
            int mDicLength = d.get(i).length();
            while (mDicIndex < mDicLength && mWordIndex < mWordLength) {
                if (mDicContent.charAt(mDicIndex) == s.charAt(mWordIndex)) {
                    mTmp += mDicContent.charAt(mDicIndex);
                    mDicIndex++;
                }
                mWordIndex++;
            }
            if (mTmp.length() != mDicLength) { // judge substring match d list or not
                mTmp = "";
            }

            if (mFinalAnser.length() < mTmp.length()) {
                mFinalAnser = mTmp;
                // System.out.println("Get!!!" + mFinalAnser);
            } else if (mFinalAnser.length() == mTmp.length()) { // same length comparison  with the smallest lexicographical order
//                System.out.println("Compare!!!");
//                for (int j = 0; j < mFinalAnser.length(); j++) {
//                    if (mFinalAnser.charAt(j) < mTmp.charAt(j)) {
//                        break;
//                    } else if (mFinalAnser.charAt(j) > mTmp.charAt(j)) {
//                        mFinalAnser = mTmp;
//                        break;
//                    }
//                }
                if (mFinalAnser.compareTo(mTmp) > 0) {
                    mFinalAnser = mTmp;
                }
            }
            // System.out.println(mFinalAnser + " " + mTmp);
        }

        return mFinalAnser;
    }
}
