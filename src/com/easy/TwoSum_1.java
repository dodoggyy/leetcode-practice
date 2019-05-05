/**
 * 
 */
package com.easy;

import java.util.HashMap;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class TwoSum_1 {

	// Runtime: 2 ms, faster than 99.47%
	// Memory Usage: 38.4 MB, less than 51.47%
	public int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> mMap = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            mMap.put(target - nums[i], i);
        }
        for (int i = 0; i < nums.length; i++) {
            if (mMap.containsKey(nums[i]) && (mMap.get(nums[i]) != i)) {
                return new int[]{ (i) , (mMap.get(nums[i]))};
            }
        }
        return null;
    }
}
