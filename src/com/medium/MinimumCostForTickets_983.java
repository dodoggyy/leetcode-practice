/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class MinimumCostForTickets_983 {
    private int[] days;
    private int[] costs;
    private Integer[] mDP;

    /**
     * @param args
     */
    public static void main(String[] args) {
        // TODO Auto-generated method stub

    }

    // 2019年8月15日
    // Use backtracking
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 1 ms, faster than 93.99%
    // Memory Usage: 35.8 MB, less than 100.00%
    public int mincostTickets(int[] days, int[] costs) {
        this.days = days;
        this.costs = costs;
        mDP = new Integer[days.length];
        int mResult = helper(0);

        return mResult;
    }

    private int helper(int mIndex) {
        if (mIndex >= days.length) {
            return 0;
        }
        if (mDP[mIndex] != null) {
            return mDP[mIndex];
        }
        int mOneDayCost = costs[0] + helper(mIndex + 1);
        int mSevenDayIndex = mIndex + 1;
        while (mSevenDayIndex < days.length && days[mSevenDayIndex] <= days[mIndex] + 6) {
            mSevenDayIndex++;
        }
        int mSevenDayCost = costs[1] + helper(mSevenDayIndex);

        int mThirtyDayIndex = mIndex + 1;
        while (mThirtyDayIndex < days.length && days[mThirtyDayIndex] <= days[mIndex] + 29) {
            mThirtyDayIndex++;
        }
        int mThirtyDayCost = costs[2] + helper(mThirtyDayIndex);

        mDP[mIndex] = Math.min(mOneDayCost, Math.min(mSevenDayCost, mThirtyDayCost));
        return mDP[mIndex];
    }
}
