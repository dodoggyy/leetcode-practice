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
public class ContainsDuplicate_217 {
    private static final int DEFAULT_VALUE = 1;

    // Use hash map
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 9 ms, faster than 53.64%
    // Memory Usage: 43.7 MB, less than 35.36%
    public boolean containsDuplicate(int[] nums) {
        Map<Integer, Integer> mMap = new HashMap<>();
        for (int num : nums) {
            if (mMap.containsKey(num)) {
                return true;
            } else {
                mMap.put(num, DEFAULT_VALUE);
            }
        }
        return false;
    }

    // Use sorting
    // Time Complexity: O(nlogn + n) = O(nlogn)
    // Space Complexity:O(1)
    // Runtime: 3 ms, faster than 99.08%
    // Memory Usage: 43.9 MB, less than 32.29%
    public boolean containsDuplicate2(int[] nums) {
        Arrays.sort(nums);
        for (int i = 0; i < nums.length - 1; i++) {
            if (nums[i] == nums[i + 1]) {
                return true;
            }
        }
        return false;
    }
}
