/**
 * 
 */
package com.medium;

import java.util.Arrays;
import java.util.PriorityQueue;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class KthLargestElementinanArray_215 {

    public static void main(String[] args) {
        int[] nums = {3,2,3,1,2,4,5,5,6};//{3,2,1,5,6,4};
        int k = 4;//2;
        
        System.out.println(findKthLargest(nums, k));
    }
    
    public static int findKthLargest(int[] nums, int k) {
        
        //return priorityQueue(nums, k);
        //return libSort(nums, k);
        return sort(nums, k);
    }

    // Runtime: 6 ms, faster than 64.98% 
    // Memory Usage: 39.7 MB, less than 5.17%
    private static int priorityQueue(int[] nums, int k) {
        PriorityQueue<Integer> mPriorityQueue = new PriorityQueue<>();
        for(int i = 0; i < nums.length; i++) {
            mPriorityQueue.add(nums[i]);
            if(mPriorityQueue.size() > k) {
                mPriorityQueue.remove(); // also can use poll()
            }
        }
        return mPriorityQueue.element(); // also can use peek()
        
    }
    
    // Runtime: 2 ms, faster than 97.64%
    // Memory Usage: 38 MB, less than 92.53%
    private static int libSort(int[] nums, int k) {
        Arrays.sort(nums); // use lib sort
        return nums[nums.length - k];
    }

    private static int sort(int[] nums, int k) {
        quickSort(nums, 0, nums.length-1);
//        for(int i = 0; i < nums.length;i++) {
//            System.out.printf("%d ",nums[i]);
//        }
        return nums[nums.length - k];
    }

    // Quick sort with DS version
    // Runtime: 28 ms, faster than 35.46%
    // Memory Usage: 39.2 MB, less than 13.73%
    private static void quickSort(int[] nums, int mIndexStart, int mIndexEnd) {
        if(mIndexStart < mIndexEnd) {

        //System.out.println("test");
        int mPivot, mTmpIndexLeft = mIndexStart, mTmpIndexRight = mIndexEnd;
        mPivot = nums[mIndexStart];
        while (mTmpIndexLeft < mTmpIndexRight) {

            while((nums[mTmpIndexLeft] <= mPivot) && (mTmpIndexLeft < mIndexEnd)) {
                mTmpIndexLeft++;
            }
            while((nums[mTmpIndexRight] > mPivot) && (mTmpIndexRight > mIndexStart)) {
                mTmpIndexRight--;
            }
            
            if (mTmpIndexLeft < mTmpIndexRight) {
//                System.out.printf("IndexLeft:%d IndexRight:%d\n",mTmpIndexLeft, mTmpIndexRight);
                swap(nums, mTmpIndexLeft, mTmpIndexRight);
            }
        }
        
        swap(nums, mIndexStart, mTmpIndexRight);
        
//        System.out.printf("mIndexStart:%d mIndexEnd:%d\n", mIndexStart, mIndexEnd);
//        System.out.printf("mTmpIndexLeft:%d mTmpIndexRight:%d\n", mTmpIndexLeft, mTmpIndexRight);
        
        quickSort(nums, mIndexStart, (mTmpIndexRight - 1));
        quickSort(nums, (mTmpIndexRight + 1), mIndexEnd);
        }
    }
    
    private static int  partition(int[] nums, int mIndexLeft, int mIndexRight) {
        int i = mIndexLeft - 1;
        for(int j = mIndexLeft; j < mIndexRight; j++) {
            if(nums[j] < nums[mIndexRight]) {
                i++;
                swap(nums, i , j);
            }
        }
        swap(nums, (i+1), mIndexRight);
        return (i+1);
    }
    
    // Quick sort with Algorithm version (partition)
    // Runtime: 84 ms, faster than 5.22%
    // Memory Usage: 38.8 MB, less than 57.34%
    private static void quickSort2(int[] nums, int mIndexLeft, int mIndexRight) {
        if(mIndexLeft < mIndexRight) {
            int q = partition(nums, mIndexLeft, mIndexRight);
            quickSort2(nums, mIndexLeft, q-1);
            quickSort2(nums, q+1, mIndexRight);
        }
    }

    private static void swap(int[] nums, int mIndex1, int mIndex2) {
//        System.out.printf("%d %d\n", mIndex1, mIndex2);
        int mTmp = nums[mIndex1];
        nums[mIndex1] = nums[mIndex2];
        nums[mIndex2] = mTmp;
    }
}
