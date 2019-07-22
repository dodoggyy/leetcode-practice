/**
 * 
 */
package com.easy;

import java.util.Arrays;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class Heaters_475 {
    // 2019年7月22日
    // Check left right side distance
    // Time Complexity: O(nlogn + mlogn) = O((n+m)*logn)
    // m: length of houses, n: length of heaters
    // Space Complexity:O(1)
    // Runtime: 13 ms, faster than 49.96%
    // Memory Usage: 39.1 MB, less than 100.00%
    public int findRadius(int[] houses, int[] heaters) {
        int mResult = 0;
        Arrays.sort(heaters);

        for (int house : houses) {
            int mIndex = Arrays.binarySearch(heaters, house);
            if (mIndex < 0) { // not found
                mIndex = ~mIndex;
                // judge left side
                int mDistance1 = mIndex - 1 >= 0 ? house - heaters[mIndex - 1] : Integer.MAX_VALUE;
                // judge right side
                int mDistance2 = mIndex < heaters.length ? heaters[mIndex] - house : Integer.MAX_VALUE;

                mResult = Math.max(mResult, Math.min(mDistance1, mDistance2));
            }
        }
        return mResult;
    }

    // 2019年7月22日
    // Use 2 pointers
    // Time Complexity: O(mlogm + nlogn + m*n)
    // Space Complexity:O(1)
    // m: length of houses, n: length of heaters
    // Runtime: 9 ms, faster than 86.96%
    // Memory Usage: 40.5 MB, less than 93.75%
    public int findRadius2(int[] houses, int[] heaters) {
        int mResult = 0;
        Arrays.sort(houses);
        Arrays.sort(heaters);

        int mIndex = 0;
        for (int house : houses) {
            while (mIndex < heaters.length - 1
                    && Math.abs(heaters[mIndex] - house) >= Math.abs(heaters[mIndex + 1] - house)) {
                mIndex++;
            }
            mResult = Math.max(mResult, Math.abs(heaters[mIndex] - house));
        }
        return mResult;
    }
}
