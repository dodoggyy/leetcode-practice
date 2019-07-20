/**
 * 
 */
package com.easy;

import java.sql.SQLException;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.Iterator;
import java.util.List;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class FindAllNumbersDisappearedInAnArray_448 {
    public static void main(String[] args) throws SQLException {
        FindAllNumbersDisappearedInAnArray_448 mTest = new FindAllNumbersDisappearedInAnArray_448();
        int nums[] = { 4, 3, 2, 7, 8, 2, 3, 1 };
        List<Integer> mResult = mTest.findDisappearedNumbers3(nums);
        Iterator<Integer> mIter = mResult.iterator();
        while (mIter.hasNext()) {
            System.out.println(mIter.next());
        }
    }

    // 2019年7月20日
    // Use HashSet
    // Time Complexity: O(2n) = O(n)
    // Space Complexity:O(n)
    // Runtime: 23 ms, faster than 22.76%
    // Memory Usage: 52.6 MB, less than 16.80%
    public List<Integer> findDisappearedNumbers(int[] nums) {
        List<Integer> mResult = new ArrayList<>();
        HashSet<Integer> mSet = new HashSet<>();
        for (int i = 0; i < nums.length; i++) {
            mSet.add(nums[i]);
        }
        for (int i = 1; i <= nums.length; i++) {
            if (!mSet.contains(i)) {
                mResult.add(i);
            }
        }

        return mResult;
    }

    // 2019年7月20日
    // Use SWAP to index
    // Time Complexity: O(2n) = O(n)
    // Space Complexity:O(1)
    // Runtime: 6 ms, faster than 80.20%
    // Memory Usage: 46.7 MB, less than 98.67%
    public List<Integer> findDisappearedNumbers2(int[] nums) {
        List<Integer> mResult = new ArrayList<>();

        for (int i = 0; i < nums.length; i++) {
            if (nums[i] != nums[nums[i] - 1]) {
                swap(nums, i, nums[i] - 1);
                i--; //
            }

        }
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] != (i + 1)) {
                mResult.add(i + 1);
            }
        }

        return mResult;
    }

    private void swap(int[] nums, int mIndex1, int mIndex2) {
        nums[mIndex1] = nums[mIndex1] ^ nums[mIndex2];
        nums[mIndex2] = nums[mIndex1] ^ nums[mIndex2];
        nums[mIndex1] = nums[mIndex1] ^ nums[mIndex2];

    }

    // 2019年7月20日
    // Use assign index value to negative to judge who is positive
    // Time Complexity: O(2n) = O(n)
    // Space Complexity:O(1)
    // Runtime: 6 ms, faster than 80.20%
    // Memory Usage: 47.8 MB, less than 83.91%
    public List<Integer> findDisappearedNumbers3(int[] nums) {
        List<Integer> mResult = new ArrayList<>();
        for (int i = 0; i < nums.length; i++) {
            int mIndex = Math.abs(nums[i]) - 1;
            if (nums[mIndex] > 0) {
                nums[mIndex] = -nums[mIndex];
            }
        }
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] > 0) {
                mResult.add(i + 1);

            }
        }

        return mResult;

    }
}
