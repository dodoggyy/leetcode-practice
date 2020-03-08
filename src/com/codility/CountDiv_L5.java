/**
 * 
 */
package com.codility;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class CountDiv_L5 {
    // 2020年3月9日
    // Use simple for loop:
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    public int solution(int A, int B, int K) {
        // write your code in Java SE 8
        int result = 0;
        for (int i = A; i <= B; i++) {
            if (i % K == 0) {
                result++;
            }
        }
        return result;
    }

    // 2020年3月9日
    // Use simple calculate:
    // Time Complexity: O(1)
    // Space Complexity:O(1)
    public int solution2(int A, int B, int K) {
        if ((B - A) < K) {
            int tmp = A % K;
            if (tmp != 0 && (A + (K - tmp) > B)) {
                return 0;
            } else {
                return 1;
            }
        }

        int start = (A % K == 0) ? A : A + (K - A % K);

        return (B - start) / K + 1;
    }
}
