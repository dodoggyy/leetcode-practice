/**
 * 
 */
package com.medium;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ThreeSum_15 {
    // 2019年8月6日
    // Reduce 2-sum problem and use 2 pointers
    // Time Complexity: O(nlogn + n^2) = O(n^2)
    // Space Complexity:O(1)
    // Runtime: 827 ms, faster than 5.00%
    // Memory Usage: 52.4 MB, less than 6.62%
    public List<List<Integer>> threeSum(int[] nums) {
        List<List<Integer>> mResult = new ArrayList<List<Integer>>();

        if (nums == null || nums.length < 3) {
            return mResult;
        }

        Arrays.sort(nums);

        for (int i = 0; i < nums.length - 2; i++) {
            int mComplement = 0 - nums[i];

            int mIndexFront = i + 1, mIndexTail = nums.length - 1;

            while (mIndexFront < mIndexTail) {
                if (nums[mIndexFront] + nums[mIndexTail] == mComplement) {
                    List<Integer> mList = Arrays.asList(nums[i], nums[mIndexFront], nums[mIndexTail]);
                    if (!mResult.contains(mList)) {
                        mResult.add(mList);
                    }
                    mIndexFront++;
                    mIndexTail--;
                } else if (nums[mIndexFront] + nums[mIndexTail] < mComplement) {
                    mIndexFront++;
                } else {
                    mIndexTail--;
                }
            }

        }

        return mResult;
    }

    // 2019年8月6日
    // Optimize reduce 2-sum problem and use 2 pointers
    // Time Complexity: O(nlogn + n^2) = O(n^2)
    // Space Complexity:O(1)
    // Runtime: 31 ms, faster than 72.23%
    // Memory Usage: 47.4 MB, less than 89.25%
    public List<List<Integer>> threeSum2(int[] nums) {
        List<List<Integer>> mResult = new ArrayList<List<Integer>>();

        if (nums == null || nums.length < 3) {
            return mResult;
        }

        Arrays.sort(nums);

        for (int i = 0; i < nums.length - 2; i++) {
            if (i > 0 && nums[i] == nums[i - 1]) {
                continue;
            }
            int mComplement = 0 - nums[i];
            int mIndexFront = i + 1;
            int mIndexTail = nums.length - 1;
            while (mIndexFront < mIndexTail) {
                if (nums[mIndexFront] + nums[mIndexTail] == mComplement) {
                    List<Integer> mList = Arrays.asList(nums[i], nums[mIndexFront], nums[mIndexTail]);
                    mResult.add(mList);

                    while (mIndexFront < mIndexTail && nums[mIndexFront] == nums[mIndexFront + 1]) {
                        mIndexFront++;
                    }
                    while (mIndexFront < mIndexTail && nums[mIndexTail] == nums[mIndexTail - 1]) {
                        mIndexTail--;
                    }
                    mIndexFront++;
                    mIndexTail--;
                } else if (nums[mIndexFront] + nums[mIndexTail] < mComplement) {
                    mIndexFront++;
                } else {
                    mIndexTail--;
                }
            }
        }

        return mResult;
    }
}
