/**
 * 
 */
package com.easy;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class MajorityElement_169 {
    // 2019年7月10日
    // Use HashMap
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 11 ms, faster than 39.33%
    // Memory Usage: 40.5 MB, less than 98.55%
    public int majorityElement(int[] nums) {
        if (nums == null || nums.length == 0) {
            return 0;
        }
        if (nums.length == 1) {
            return nums[0];
        }
        Map<Integer, Integer> mMap = new HashMap<>();
        for (int i = 0; i <= nums.length; i++) {
            mMap.put(nums[i], mMap.getOrDefault(nums[i], 0) + 1);
        }
        for (int key : mMap.keySet()) {
            if (mMap.get(key) > nums.length / 2) {
                return key;
            }
        }

        return 0;
    }

    // 2019年7月10日
    // Use sorting
    // Time Complexity: O(nlogn)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 42 MB, less than 62.56%
    public int majorityElement2(int[] nums) {
        Arrays.sort(nums);
        return nums[nums.length / 2];
    }
}
