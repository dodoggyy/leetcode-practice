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

    // 2020年2月1日
    // Use Boyer-Moore vote
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 99.84%
    // Memory Usage: 41.8 MB, less than 68.38%
    public int majorityElement3(int[] nums) {
        int count = 0;
        int mResult = nums[0];

        for (int num : nums) {
            if (mResult == num) {
                count++;
            } else {
                count--;
                if (count == 0) {
                    count = 1;
                    mResult = num;
                }
            }
        }

        return mResult;
    }

    // 2020年2月1日
    // Use Bit vote
    // Time Complexity: O(32*n)
    // Space Complexity:O(1)
    // Runtime: 4 ms, faster than 53.62%
    // Memory Usage: 41.5 MB, less than 71.32%
    public int majorityElement4(int[] nums) {
        int size = nums.length;
        int mResult = 0;

        for (int i = 0; i < 32; i++) {
            int count = 0;
            int mask = 1 << i;
            for (int num : nums) {
                if ((num & mask) != 0) {
                    count++;
                }
            }
            if (count > size / 2) {
                mResult |= mask;
            }

        }
        return mResult;
    }
}
