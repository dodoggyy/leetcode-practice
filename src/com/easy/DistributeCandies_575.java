/**
 * 
 */
package com.easy;

import java.util.HashSet;
import java.util.Set;

/**
 * @author Chris Lin
 * @version 1.0
 * 
 */
public class DistributeCandies_575 {
    // Use hash set
    // Time Complexity: O(n)
    // Space Complexity: O(n)
    // Runtime: 27 ms, faster than 95.21%
    // Memory Usage: 40 MB, less than 99.28%
    public int distributeCandies(int[] candies) {
        Set<Integer> mSet = new HashSet<>();
        int mTarget = candies.length / 2;
        for (int candy : candies) {
            mSet.add(candy);
            if (mSet.size() >= mTarget) {
                break;
            }
        }
        return mSet.size();
    }

}
