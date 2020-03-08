/**
 * 
 */
package com.codility;

import java.util.Arrays;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class Triangle_L6 {
    // 2020年3月9日
    // Use sort:
    // Time Complexity: O(nlogn)
    // Space Complexity:O(1)
    public int solution(int[] A) {
        // write your code in Java SE 8
        if(A == null || A.length < 3) {
            return 0;
        }
        Arrays.sort(A);
        for (int i = 0; i < A.length - 2; i++) {
            if(A[i] >= 0 && A[i] > A[i+2] - A[i+1]) {
                return 1;
            }
        }
        
        return 0;
    }
}
