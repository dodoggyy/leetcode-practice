/**
 * 
 */
package com.medium;

import java.util.HashMap;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
/**
 * @author Chris Lin
 * @version 1.0
 */
public class SortCharactersByFrequency_451 {
    public static void main(String[] args) {
        String s = "BbAabb";
        frequencySort(s);
    }

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
                for (char c : mFrequencyBucket[i]) { // list all of char in bucket
                    for (int j = 0; j < i; j++) { // add char according bucket number
                        mStrAnswer.append(c);
                    }
                }
            }

        }
        // System.out.println(mAnswer);

        return mStrAnswer.toString();
    }
}
