/**
 * 
 */
package com.easy;

import java.util.HashMap;
import java.util.Map;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ContainsDuplicateII_219 {
    // Use Hash Map
    // Time Complexity: O(n)
    // Space Complexity: O(n)
    // Runtime: 9 ms, faster than 47.62%
    // Memory Usage: 41.7 MB, less than 96.85%
    public boolean containsNearbyDuplicate(int[] nums, int k) {
        Map<Integer, Integer> mMap = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            if (mMap.containsKey(nums[i]) && (i - mMap.get(nums[i]) <= k)) {
                return true;
            } else {
                mMap.put(nums[i], i);
            }

        }
        return false;
    }
}
