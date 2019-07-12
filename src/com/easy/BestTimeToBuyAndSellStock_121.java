/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class BestTimeToBuyAndSellStock_121 {
    // 2019年7月10日
    // Use brute force
    // Time Complexity: O(n^2)
    // Space Complexity:O(1)
    // Runtime: 256 ms, faster than 5.02%
    // Memory Usage: 37.6 MB, less than 99.48%
    public int maxProfit(int[] prices) {
        int mLength = prices.length;
        int mMaxProfit = 0;
        for (int i = 0; i < mLength - 1; i++) {
            for (int j = i; j < mLength; j++) {
                if (prices[j] - prices[i] > mMaxProfit) {
                    mMaxProfit = prices[j] - prices[i];
                }
            }

        }
        return mMaxProfit;
    }
    
    // 2019年7月10日
    // One pass
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 88.01%
    // Memory Usage: 37.9 MB, less than 98.45%
    public int maxProfit2(int[] prices) {
        int mLength = prices.length;
        int mMaxProfit = 0, mMin = Integer.MAX_VALUE;
        for (int i = 0; i < mLength; i++) {
            if(prices[i] < mMin) {
                mMin = prices[i];
            } else if(prices[i] - mMin > mMaxProfit){
                mMaxProfit = prices[i] - mMin;
            }
        }
        return mMaxProfit;
    }
}
