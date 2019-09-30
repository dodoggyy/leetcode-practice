/**
 * 
 */
package com.medium;

import java.util.HashMap;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.util.PriorityQueue;

/**
 * @author Chris Lin
 * @version 1.1
 */
public class SortCharactersByFrequency_451 {
    public static void main(String[] args) {
        String s = "BbAabb";
        frequencySort(s);
    }

    // Bucket Sort
    // Time complexity: O(n)
    // Space complexity: O(n)
    // Runtime: 14 ms, faster than 91.77%
    // Memory Usage: 38.7 MB, less than 97.17%
    public static String frequencySort(String s) {
        Map<Character, Integer> mMap = new HashMap<Character, Integer>();
        // for(int i = 0; i < mAnswer.length; i++) {
        // mMap.put(mAnswer[i], mMap.getOrDefault(mMap.get(mAnswer[i]), 0)+1);
        // }
        for (char c : s.toCharArray()) {
            mMap.put(c, mMap.getOrDefault(c, 0) + 1);
        }

        List<Character>[] mFrequencyBucket = new ArrayList[s.length() + 1];
        for (char c : mMap.keySet()) {
            int mCount = mMap.get(c);
            if (mFrequencyBucket[mCount] == null) {
                mFrequencyBucket[mCount] = new ArrayList<Character>();
            }
            mFrequencyBucket[mCount].add(c);
        }
        StringBuilder mStrAnswer = new StringBuilder();
        for (int i = mFrequencyBucket.length - 1; i >= 0; i--) {
            if (mFrequencyBucket[i] != null) {
                for (char c : mFrequencyBucket[i]) { // list all of char in
                                                     // bucket
                    for (int j = 0; j < i; j++) { // add char according bucket
                                                  // number
                        mStrAnswer.append(c);
                    }
                }
            }

        }
        // System.out.println(mAnswer);

        return mStrAnswer.toString();
    }

    // Bucket Sort 2
    // Time complexity: O(n)
    // Space complexity: O(n)
    // Runtime: 17 ms, faster than 73.52%
    // Memory Usage: 39 MB, less than 85.19%
    public String frequencySort2(String s) {
        HashMap<Character, Integer> mMap = new HashMap<>();
        List<Character>[] mBucket = new List[s.length() + 1];
        StringBuilder mResult = new StringBuilder();
        for (char c : s.toCharArray()) {
            mMap.put(c, mMap.getOrDefault(c, 0) + 1);
        }
        for (char c : mMap.keySet()) {
            int mFreq = mMap.get(c);
            if (mBucket[mFreq] == null) {
                mBucket[mFreq] = new ArrayList<>();
            }
            mBucket[mFreq].add(c);
        }
        for (int i = mBucket.length - 1; i >= 0; i--) {
            if (mBucket[i] != null) {
                for (char c : mBucket[i]) {
                    for (int j = 0; j < i; j++) {
                        mResult.append(c);
                    }
                }
            }
        }

        return mResult.toString();
    }

    // Use Priority Queue
    // Time complexity: O(nlogm) = O(n)
    // Space complexity: O(n)
    // m: 26 characters
    // Runtime: 44 ms, faster than 41.91%
    // Memory Usage: 37.9 MB, less than 96.30%
    public String frequencySort3(String s) {
        HashMap<Character, Integer> mMap = new HashMap<>();
        for (char c : s.toCharArray()) {
            mMap.put(c, mMap.getOrDefault(c, 0) + 1);
        }
        PriorityQueue<Character> mQueue = new PriorityQueue<>((x, y) -> mMap.get(y) - mMap.get(x));
        for (char c : mMap.keySet()) {
            mQueue.add(c);
        }
        StringBuilder mResult = new StringBuilder();
        while (!mQueue.isEmpty()) {
            char c = mQueue.poll();
            int mFreq = mMap.get(c);
            for (int i = 0; i < mFreq; i++) {
                mResult.append(c);
            }
        }
        return mResult.toString();
    }
}
