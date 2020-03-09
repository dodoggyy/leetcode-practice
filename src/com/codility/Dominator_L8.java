/**
 * 
 */
package com.codility;

import java.util.HashMap;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class Dominator_L8 {
    // 2020年3月9日
    // Use hash map:
    // Time Complexity: O(nlogn)
    // Space Complexity:O(n)
    public int solution(int[] A) {
        if (A == null || A.length == 1) {
            return 0;
        }
        HashMap<Integer, Integer> map = new HashMap<>();
        int size = A.length;
        for (int i = 0; i < size; i++) {
            if (map.containsKey(A[i])) {
                int count = map.get(A[i]) + 1;
                if (count > size / 2) {
                    return i;
                }
                map.put(A[i], count);
            } else {
                map.put(A[i], 1);
            }
        }
        return -1;
    }
}
