/**
 * 
 */
package com.medium;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class SortColors_75 {

    /**
     * @param args
     */
    public static void main(String[] args) {
        // TODO Auto-generated method stub
        int[] nums = { 2, 0, 1 };
        sortColors2(nums);
        for (Integer num : nums) {
            System.out.println(num);
        }
    }

    // Counting Sort
    // Time complexity: O(2n) = O(n)
    // Space complexity: O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 37.3 MB, less than 39.77%
    public void sortColors(int[] nums) {
        int[] mCount = new int[3];
        int mIndex = 0;
        for (int i = 0; i < nums.length; i++) {
            mCount[nums[i]]++;
        }
        for (int i = 0; i < 3; i++) {
            for (int j = mCount[i]; j > 0; j--) {
                nums[mIndex++] = i;
            }
        }
    }

    // one-pass algorithm
    // Time complexity: O(n)
    // Space complexity: O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 37.3 MB, less than 52.02%
    public static void sortColors2(int[] nums) {
        int mIndexRed = 0, mIndexBlue = nums.length - 1;
        // for(int i = 0; i <= mIndexBlue;++i) {
        // if(nums[i] == 2) { // blue case
        // swap(nums, mIndexBlue--, i--);
        // } else if(nums[i] == 0) { // red case
        // swap(nums, mIndexRed++, i);
        // }
        // }
        int mIndex = 0;
        while (mIndex <= mIndexBlue) {
            switch (nums[mIndex]) {
            case 0:
                swap(nums, mIndexRed++, mIndex++);
                break;
            case 1:
                mIndex++;
                break;
            case 2:
                swap(nums, mIndexBlue--, mIndex);
                break;
            default:
                System.out.println("input error");
                break;
            }
        }
    }

    public static void swap(int[] nums, int mIndex1, int mIndex2) {
        int tmp = nums[mIndex1];
        nums[mIndex1] = nums[mIndex2];
        nums[mIndex2] = tmp;
    }

    // one-pass algorithm 2
    // Time complexity: O(n)
    // Space complexity: O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 35.4 MB, less than 99.21%
    public void sortColors3(int[] nums) {
        int mIndexRed = -1, mIndexWhite = -1, mIndexBlue = -1;
        for (int i = 0; i < nums.length; i++) {
            switch (nums[i]) {
            case 0:
                nums[++mIndexBlue] = 2;
                nums[++mIndexWhite] = 1;
                nums[++mIndexRed] = 0;
                break;
            case 1:
                nums[++mIndexBlue] = 2;
                nums[++mIndexWhite] = 1;
                break;
            case 2:
                nums[++mIndexBlue] = 2;
                break;
            }
        }
    }
}
