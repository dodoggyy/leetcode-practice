/**
 * 
 */
package com.codility;


/**
 * @author Chris Lin
 * @version 1.0
 */
public class CountFactors_L10 {
    // 2020年3月10日
    // Use math calculation
    // Time Complexity: O(n^1/2)
    // Space Complexity:O(1)
    public int solution(int N) {
        // write your code in Java SE 8
        int result = 0;
        for (int i = 1; i * i <= N; i++) {
            if (i * i == N) {
                result++;
                continue;
            }
            if (N % i == 0) {
                result += 2;
            }
        }
        return result;
    }
}
