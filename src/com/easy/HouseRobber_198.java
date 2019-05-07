/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class HouseRobber_198 {

    // DP[i] = max(DP[i-2] + nums[i], DP[i-1])
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 33.5 MB, less than 100.00%
    public int rob(int[] nums) {
        int mRob = 0;
        int mNotRob = 0;
        for (int i = 0; i < nums.length; i++) {
            int mCurrent = Math.max(mNotRob + nums[i], mRob);
            mNotRob = mRob;
            mRob = mCurrent;
        }

        return mRob;
    }
    
    // Judge even & odd position
    // Runtime: 0 ms, faster than 100.00% 
    // Memory Usage: 33.2 MB, less than 100.00%
    public int rob2(int[] nums) {
        int mEvenRob = 0, mOddRob = 0;
        for(int i = 0; i < nums.length; i++) {
            if((i%2) == 0) {
                mEvenRob = Math.max(mEvenRob + nums[i], mOddRob);
            } else {
                mOddRob = Math.max(mOddRob + nums[i], mEvenRob);
            }
        }
        
        return Math.max(mEvenRob, mOddRob);
    }
}
