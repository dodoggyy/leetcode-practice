/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class FindSmallestLetterGreaterThanTarget_744 {
    // 2019年10月2日
    // Use linear search
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 38.7 MB, less than 100.00%
    public char nextGreatestLetter(char[] letters, char target) {
        for (char c : letters) {
            if (c - target > 0) {
                return c;
            }
        }
        return letters[0];
    }

    // 2019年10月2日
    // Use binary search
    // Time Complexity: O(logn)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 38.7 MB, less than 100.00%
    public char nextGreatestLetter2(char[] letters, char target) {
        // mIndexHigh starts at 'letters.length' rather than the usual
        // 'letters.length - 1'.
        // It is because the terminal condition is 'mIndexLow < mIndexHigh' and
        // if mIndexHigh starts from 'letters.length - 1',
        // we can never consider value at index 'letters.length - 1'
        int mLength = letters.length;
        if (target >= letters[mLength - 1]) {
            return letters[0];
        }
        int mIndexLow = 0, mIndexHigh = mLength - 1;

        // Terminal condition is 'mIndexLow < mIndexHigh', to avoid infinite
        // loop when target is smaller than the first element
        while (mIndexLow < mIndexHigh) {
            int mIndexMid = mIndexLow + (mIndexHigh - mIndexLow) / 2;
            int mCompare = letters[mIndexMid] - target;
            if (mCompare > 0) {
                mIndexHigh = mIndexMid;
            } else {
                mIndexLow = mIndexMid + 1;
            }
        }

        // Because lo can end up pointing to index 'letters.length', in which
        // case we return the first element
        return letters[mIndexHigh];
    }
}
