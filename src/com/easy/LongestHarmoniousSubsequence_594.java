/**
 * 
 */
package com.easy;

import java.util.HashMap;
import java.util.Iterator;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class LongestHarmoniousSubsequence_594 {
    // 2019年7月19日
    // Use Recursive
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 32 ms, faster than 46.02%
    // Memory Usage: 40 MB, less than 97.06%
    public int findLHS(int[] nums) {
        int mResult = 0;
        HashMap<Integer, Integer> mMap = new HashMap<>();
        for (int num : nums) {
            mMap.put(num, mMap.getOrDefault(num, 0) + 1);
        }
        Iterator<Integer> mIter = mMap.keySet().iterator();
        while (mIter.hasNext()) {
            int mTmpNum = (int) mIter.next();
            int mTmpNumCount = mMap.get(mTmpNum);
            if (mMap.containsKey(mTmpNum + 1)) {
                mResult = Math.max(mResult, mTmpNumCount + mMap.get(mTmpNum + 1));
            }

        }
        return mResult;
    }
}
