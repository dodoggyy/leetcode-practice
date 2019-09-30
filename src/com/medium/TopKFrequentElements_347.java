/**
 * 
 */
package com.medium;

import java.util.ArrayList;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.PriorityQueue;

/**
 * @author Chris Lin
 * @version 1.1
 */
public class TopKFrequentElements_347 {
    
    // Priority queue / max heap
    // Time complexity: O(n) + O(nlogk)
    // Space complexity: O(n)
    // Runtime: 45 ms, faster than 31.60% 
    // Memory Usage: 37.7 MB, less than 96.85%
    public List<Integer> topKFrequent(int[] nums, int k) {
        Map<Integer, Integer> mHashMap = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
//            if(mHashMap.containsKey(nums[i])) {
//                mHashMap.replace(nums[i], mHashMap.get(nums[i])+1);
//            } else {
//                mHashMap.put(nums[i], 1);
//            }
            mHashMap.put(nums[i], mHashMap.getOrDefault(nums[i], 0) + 1);
        }
        
//        for(Integer key : mHashMap.keySet()) {
//            System.out.println(key + " : " + mHashMap.get(key));
//        }
        
        PriorityQueue<int[]> mPriorityQueue = new PriorityQueue<>((x,y) -> x[0] - y[0]);
        for(Integer num : mHashMap.keySet()) {
            // add() fail may throw exception while offer() may return false
            mPriorityQueue.offer(new int[] {mHashMap.get(num), num});
            
            if(mPriorityQueue.size() > k) {
                mPriorityQueue.poll(); // remove() fail may throw exception while poll() may return null
            }
        }
        
        List<Integer> mArrayListAnswer = new ArrayList<Integer>();
        for(int i = 0; i < k; i++) {
            mArrayListAnswer.add(mPriorityQueue.poll()[0]);
        }
        Collections.reverse(mArrayListAnswer);
        
        return mArrayListAnswer;

    }
    
    // Bucket Sort
    // Time complexity: O(n)
    // Space complexity: O(n)
    // Runtime: 11 ms, faster than 93.03%
    // Memory Usage: 40.9 MB, less than 62.80%
    public List<Integer> topKFrequent2(int[] nums, int k) {
        Map<Integer, Integer> mHashMap = new HashMap<>();
        List<Integer>[] mBucket = new ArrayList[nums.length+1];
        for (int i = 0; i < nums.length; i++) {
            mHashMap.put(nums[i], mHashMap.getOrDefault(nums[i], 0) + 1);
        }
        
        for (Integer num : mHashMap.keySet()) {
            int mCount = mHashMap.get(num);
            if (mBucket[mCount] == null) {
                mBucket[mCount] = new ArrayList<>();
            }
            mBucket[mCount].add(num);
        }
        
        List mArrayListAnswer = new ArrayList<>();
        for (int i = (mBucket.length-1); i > 0 && mArrayListAnswer.size() < k; i--) {
            if(mBucket[i] != null) {
                mArrayListAnswer.addAll(mBucket[i]);
            }
        }
        
        return mArrayListAnswer;
    }
    
    // Priority queue / max heap 2
    // Time complexity: O(n) + O(nlogk) = O(nlogk)
    // Space complexity: O(n)
    // Runtime: 48 ms, faster than 13.86%
    // Memory Usage: 39.6 MB, less than 73.28%
    public List<Integer> topKFrequent3(int[] nums, int k) {
        List<Integer> mResult = new ArrayList<>();
        HashMap<Integer, Integer> mMap = new HashMap<>();
        for(int num:nums) {
            mMap.put(num, mMap.getOrDefault(num, 0) + 1);
        }
        
        PriorityQueue<Integer> mQueue = new PriorityQueue<Integer>((x,y) -> mMap.get(x) - mMap.get(y));
        for(int key:mMap.keySet()) {
            mQueue.add(key);
            if(mQueue.size() > k) {
                mQueue.poll();
            }
        }
        while(!mQueue.isEmpty()) {
            mResult.add(mQueue.poll());
        }
        Collections.reverse(mResult);
        return mResult;
    }
}
