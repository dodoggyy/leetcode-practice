/**
 * 
 */
package com.medium;

import java.util.Arrays;
import java.util.Comparator;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class NonOverlappingIntervals_435 {
    // 2019年9月5日
    // Use greedy algorithm
    // Time Complexity: O(nlogn + n) = O(nlogn)
    // Space Complexity:O(1)
    // Runtime: 3 ms, faster than 85.23%
    // Memory Usage: 37 MB, less than 62.50%
    public int eraseOverlapIntervals(int[][] intervals) {
        if (intervals.length == 0) {
            return 0;
        }
        Arrays.sort(intervals, new Comparator<int[]>() {
            @Override
            public int compare(int[] o1, int[] o2) {
                // TODO Auto-generated method stub
                return o1[1] - o2[1];
            }
        });
        int mCount = 0;
        int mEnd = Integer.MIN_VALUE;
        for (int i = 0; i < intervals.length; i++) {
            if (intervals[i][0] >= mEnd) {
                mEnd = intervals[i][1];
            } else {
                mCount++;
            }
        }

        return mCount;
    }
}
