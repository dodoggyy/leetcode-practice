/**
 * 
 */
package com.codility;

import java.util.Arrays;

/**
 * @author Chris Lin
 * @version 1.0
 * https://app.codility.com/demo/results/training38GDBS-5WF/
 */
public class MaxProductOfThree_L6 {

    /**
     * @param args
     */
    public static void main(String[] args) {
        // TODO Auto-generated method stub

    }

    // 2019年8月15日
    // Use simple compare
    // Time Complexity: O(nlogn)
    // Space Complexity:O(1)
    public int solution(int[] A) {
        // write your code in Java SE 8
        int mLength = A.length;
        Arrays.sort(A);

        return Math.max(A[0] * A[1] * A[mLength - 1], A[mLength - 1] * A[mLength - 2] * A[mLength - 3]);
    }
}
