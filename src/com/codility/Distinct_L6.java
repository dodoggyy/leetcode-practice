/**
 * 
 */
package com.codility;

import java.util.HashSet;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class Distinct_L6 {
    // 2020年3月9日
    // Use hash set:
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    public int solution(int[] A) {
        // write your code in Java SE 8
        HashSet<Integer> set = new HashSet<>();
        for(int num:A) {
            set.add(num);
        }
        
        return set.size();
    }
}
