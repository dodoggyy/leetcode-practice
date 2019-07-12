/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class BestTimeToBuyAndSellStockII_122 {
    public static void main(String[] args) {
        int[] mBackTest = new int[] { 2, 1, 2, 0, 1 };// {1,2,3,4,5};//{7, 1, 5,
                                                      // 3, 6, 4};
        BestTimeToBuyAndSellStockII_122 mTest = new BestTimeToBuyAndSellStockII_122();
        System.out.println(mTest.maxProfit(mBackTest));
    }

    // 2019年7月10日
    // Use Peak Valley Approach
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 94.93%
    // Memory Usage: 36.6 MB, less than 99.97%
    public int maxProfit(int[] prices) {
        if (prices == null || prices.length <= 1) {
            return 0;
        }

        int i = 0;
        int mValley = prices[0];
        int mPeak = prices[0];
        int mMaxProfit = 0;
        while (i < prices.length - 1) {
            while (i < prices.length - 1 && prices[i] >= prices[i + 1]) {
                i++;
            }
            mValley = prices[i];
            while (i < prices.length - 1 && prices[i] <= prices[i + 1]) {
                i++;
            }
            mPeak = prices[i];
            mMaxProfit += mPeak - mValley;
        }
        return mMaxProfit;
    }

    // 2019年7月10日
    // Use Simple One Pass
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 94.93%
    // Memory Usage: 36.9 MB, less than 99.97%
    public int maxProfit2(int[] prices) {
        if (prices == null || prices.length <= 1) {
            return 0;
        }
        int mMaxProfit = 0;
        for (int i = 1; i < prices.length; i++) {
            if (prices[i] - prices[i - 1] > 0) {
                mMaxProfit += prices[i] - prices[i - 1];
            }
        }
        return mMaxProfit;
    }
}
