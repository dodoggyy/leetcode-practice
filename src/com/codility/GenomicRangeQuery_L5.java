/**
 * 
 */
package com.codility;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class GenomicRangeQuery_L5 {

    public static void main(String[] args) {
        GenomicRangeQuery_L5 mTest = new GenomicRangeQuery_L5();
        String S = "CAGCCTA";
        int[] P = { 2, 5, 0 };
        int[] Q = { 4, 5, 6 };
        int[] result = new int[P.length];
        result = mTest.solution(S, P, Q);

        for (int i = 0; i < result.length; i++) {
            System.out.println(result[i]);
        }
    }

    // 2020年3月9日
    // Use counting sort:
    // Time Complexity: O(n + m)
    // Space Complexity:O(n)
    public int[] solution(String S, int[] P, int[] Q) {
        // write your code in Java SE 8
        int[][] mem = new int[S.length()][4];
        int size = P.length;
        int[] result = new int[size];
        for (int i = 0; i < S.length(); i++) {
            if (i >= 1) {
                for (int j = 0; j < 4; j++) {
                    mem[i][j] = mem[i - 1][j];
                }
            }

            switch (S.charAt(i)) {
            case 'A':
                mem[i][0]++;
                break;
            case 'C':
                mem[i][1]++;
                break;
            case 'G':
                mem[i][2]++;
                break;
            case 'T':
                mem[i][3]++;
                break;
            default:
                break;
            }
        }

//        for (int i = 0; i < S.length(); i++) {
//            for (int j = 0; j < 4; j++) {
//                System.out.print(mem[i][j] + "\t");
//            }
//            System.out.println();
//        }

        for (int i = 0; i < size; i++) {
            int start = P[i], end = Q[i];

            int maxElement = 0;

            for (int j = 0; j < 4; j++) {
                int tmp = (start == 0) ? mem[end][j] : mem[end][j] - mem[start - 1][j];
                if (tmp > 0) {
                    maxElement = j + 1;
                    break;
                }
            }

            result[i] = maxElement;
        }

        return result;
    }
}
