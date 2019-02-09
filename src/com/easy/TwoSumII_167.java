/**
 * 
 */
package com.easy;

import java.util.HashMap;
import java.util.Map;

/**
 * @author Chris
 *
 */

public class TwoSumII_167 {

    // Runtime: 3 ms, faster than 24.89% of Java online submissions for Two Sum II - Input array is sorted.
    // Memory Usage: 23.2 MB, less than 65.48% of Java online submissions for Two Sum II - Input array is sorted.
    public int[] twoSum(int[] numbers, int target) {
        int[] mResult = {1,2};//new int[2];
        HashMap<Integer, Integer> mMap = new HashMap<>();
        for (int i = 0; i < numbers.length; i++) {
            mMap.put(target - numbers[i], i);
        }
        for (int i = 0; i < numbers.length; i++) {
            if (mMap.containsKey(numbers[i]) && (mMap.get(numbers[i]) != i)) {
                mResult[1] = mMap.get(numbers[i])+1;
                mResult[0] = i+1;
                return mResult;
            }
        }
        return null;
    }
}
