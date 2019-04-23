/**
 * 
 */
package com.easy;

import java.util.HashMap;

/**
 * @author Chris Lin
 * @version 1.1
 */

public class TwoSumII_167 {

    // Runtime: 3 ms, faster than 24.89% of Java online submissions for Two Sum II - Input array is sorted.
    // Memory Usage: 23.2 MB, less than 65.48% of Java online submissions for Two Sum II - Input array is sorted.
    public int[] twoSum(int[] numbers, int target) {
        HashMap<Integer, Integer> mMap = new HashMap<>();
        for (int i = 0; i < numbers.length; i++) {
            mMap.put(target - numbers[i], i);
        }
        for (int i = 0; i < numbers.length; i++) {
            if (mMap.containsKey(numbers[i]) && (mMap.get(numbers[i]) != i)) {
                return new int[]{ (i+1) , (mMap.get(numbers[i])+1)};
            }
        }
        return null;
    }
   
    // Runtime: 0 ms, faster than 100.00% of Java online submissions for Two Sum II - Input array is sorted.
    // Memory Usage: 27.5 MB, less than 20.31% of Java online submissions for Two Sum II - Input array is sorted.
    public int[] twoSum2(int[] numbers, int target) {
        int mIndex1 = 0;
        int mIndex2 = numbers.length - 1;
        int mResult = 0;

        while(mIndex1 < mIndex2) {
            mResult = numbers[mIndex1] + numbers[mIndex2];
            if(target == mResult) {
                return new int[] {mIndex1+1, mIndex2+1};
            } else if (target < mResult) {
                mIndex2--;
            } else { // target > mResult
                mIndex1++;
            }
        }

        return null;
    }
}
