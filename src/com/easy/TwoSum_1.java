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

    // Use hash map
    // Time Complexity: O(2n) = O(n)
    // Space Complexity:O(n)
    // Runtime: 2 ms, faster than 99.47%
    // Memory Usage: 38.4 MB, less than 51.47%
    public int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> mMap = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            mMap.put(target - nums[i], i);
        }
        for (int i = 0; i < nums.length; i++) {
            if (mMap.containsKey(nums[i]) && (mMap.get(nums[i]) != i)) {
                return new int[] { (i), (mMap.get(nums[i])) };
            }
        }
        return null;
    }
    
    // Use exhaustive search
    // Time Complexity: O(n^2)
    // Space Complexity:O(1)
    // Runtime: 30 ms, faster than 20.65%
    // Memory Usage: 37.5 MB, less than 98.48%
    public int[] twoSum2(int[] nums, int target) {
        for(int i = 0; i < nums.length;i++) {
            for(int j = 1; j <nums.length; j++) {
                if(target == (nums[i] + nums[j]) && (i != j)) {
                    return new int [] {i,j};
                }
            }
        }
        return null;
    }
}
