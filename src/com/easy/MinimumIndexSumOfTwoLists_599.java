/**
 * 
 */
package com.easy;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class MinimumIndexSumOfTwoLists_599 {
    // 2019年8月3日
    // Use hash map
    // Time Complexity: O(m + n)
    // Space Complexity:O(m * x)
    // m: length of l1, n : length of l2, x: length of average string
    // Runtime: 6 ms, faster than 99.50%
    // Memory Usage: 38.7 MB, less than 99.18%
    public String[] findRestaurant(String[] list1, String[] list2) {
        HashMap<String, Integer> mMap = new HashMap<>();
        for (int i = 0; i < list1.length; i++) {
            mMap.put(list1[i], i);
        }
        List<String> mResult = new ArrayList<>();
        int mMinSum = Integer.MAX_VALUE, mSum = 0;
        for (int i = 0; i < list2.length; i++) {
            if (mMap.containsKey(list2[i])) {
                mSum = i + mMap.get(list2[i]);
                if (mSum < mMinSum) {
                    mResult.clear();
                    mResult.add(list2[i]);
                    mMinSum = mSum;
                } else if (mSum == mMinSum) {
                    mResult.add(list2[i]);
                }
            }
        }
        return mResult.toArray(new String[mResult.size()]);

    }
}
