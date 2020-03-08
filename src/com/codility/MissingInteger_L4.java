/**
 * 
 */
package com.codility;

import java.util.HashSet;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class MissingInteger_L4 {
    // 2020年3月8日
    // Use Hash Set:
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    public int solution(int[] A) {
        // write your code in Java SE 8
        HashSet<Integer> set = new HashSet<>();
        for (int num : A) {
            set.add(num);
        }

        for (int i = 1; i <= A.length; i++) {
            if (!set.contains(i)) {
                return i;
            }
        }

        return A.length + 1;
    }
}
