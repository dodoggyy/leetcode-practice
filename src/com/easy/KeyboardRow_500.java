/**
 * 
 */
package com.easy;

import java.util.Arrays;
import java.util.List;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class KeyboardRow_500 {
    // 2019年7月24日
    // Use character compare
    // Time Complexity: O(m*n)
    // Space Complexity:O(m*n)
    // m: words length, n: word length
    // Runtime: 1 ms, faster than 74.81%
    // Memory Usage: 34.4 MB, less than 99.45%
    public String[] findWords(String[] words) {
        String[] mResult = new String[words.length];
        int mLength = 0;
        Character[] mRow1 = new Character[] { 'Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P' };
        Character[] mRow2 = new Character[] { 'A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L' };
        Character[] mRow3 = new Character[] { 'Z', 'X', 'C', 'V', 'B', 'N', 'M' };
        List<Character> mList1 = Arrays.asList(mRow1);
        List<Character> mList2 = Arrays.asList(mRow2);
        List<Character> mList3 = Arrays.asList(mRow3);

        for (int i = 0; i < words.length; i++) {
            int mRecord = 0;
            boolean bIsSameRow = true;
            String mTmpStr = words[i].toUpperCase();
            for (int j = 0; j < words[i].length(); j++) {
                if (j == 0) {
                    if (mList1.contains(mTmpStr.charAt(j))) {
                        mRecord = 1;
                    } else if (mList2.contains(mTmpStr.charAt(j))) {
                        mRecord = 2;
                    } else {
                        mRecord = 3;
                    }
                } else {
                    if (mRecord == 1) {
                        if (!mList1.contains(mTmpStr.charAt(j))) {
                            bIsSameRow = false;
                            break;
                        }
                    } else if (mRecord == 2) {
                        if (!mList2.contains(mTmpStr.charAt(j))) {
                            bIsSameRow = false;
                            break;
                        }
                    } else if (mRecord == 3) {
                        if (!mList3.contains(mTmpStr.charAt(j))) {
                            bIsSameRow = false;
                            break;
                        }
                    }
                }
            }
            if (bIsSameRow) {
                mResult[mLength++] = words[i];
            }
        }

        return Arrays.copyOf(mResult, mLength);
    }

    // 2019年7月24日
    // Optimize character compare
    // Time Complexity: O(m*n)
    // Space Complexity:O(m*n)
    // m: words length, n: word length
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 34.7 MB, less than 98.07%
    public String[] findWords2(String[] words) {
        String[] mResult = new String[words.length];
        int mLength = 0;

        int[] mRows = { 2, 3, 3, 2, 1, 2, 2, 2, 1, 2, 2, 2, 3, 3, 1, 1, 1, 1, 2, 1, 1, 3, 1, 3, 1, 3 };
        char[] mCharArray;
        boolean bIsSameRow = true;
        int mRow;

        for (int i = 0; i < words.length; i++) {
            bIsSameRow = true;
            mCharArray = words[i].toUpperCase().toCharArray();
            mRow = mRows[mCharArray[0] - 'A'];

            for (int j = 1; j < mCharArray.length; j++) {
                if (mRows[mCharArray[j] - 'A'] != mRow) {
                    bIsSameRow = false;
                    break;
                }
            }

            if (bIsSameRow) {
                mResult[mLength++] = words[i];
            }
        }

        return Arrays.copyOf(mResult, mLength);
    }
}
