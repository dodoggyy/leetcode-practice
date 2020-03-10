/**
 * 
 */
package com.codility;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class WellStructure_Rakuten {

    /**
     * @param args
     */
    public static void main(String[] args) {
        // TODO Auto-generated method stub
        String S = "YAMAR";
        WellStructure_Rakuten mTest = new WellStructure_Rakuten();
        System.out.println(mTest.solution(S));
    }

    // 2020年3月10日
    // Use math permutation:
    // Time Complexity: O(m!n!)
    // Space Complexity:O(1)
    public int solution(String S) {
        // write your code in Java SE 8
        if (S == null || S.length() == 0) {
            return 0;
        }
        int[] count = new int[26];
        int size = S.length();
        for (char c : S.toCharArray()) {
            count[c - 'A']++;
        }
        int countVowel = count['A' - 'A'] + count['E' - 'A'] + count['I' - 'A'] + count['O' - 'A'] + +count['U' - 'A'];
        int countConsonant = size - countVowel;

        int diff = countConsonant - countVowel;
        if (diff > 1 || diff < 0) {
            return 0;
        }

        int[] arrVowel = new int[5];
        int[] arrConsonant = new int[21];
        int indexarrConsonant = 0, indexarrVowel = 0;

        System.out.println("countConsonant:" + countConsonant);
        System.out.println("countVowel:" + countVowel);

        for (int i = 0; i < 26; i++) {

            char c = (char) ('A' + i);
            if (c != 'A' && c != 'E' && c != 'I' && c != 'O' && c != 'U') {
                arrConsonant[indexarrConsonant++] = count[i];
            } else {
                arrVowel[indexarrVowel++] = count[i];
            }
        }

        return helper(countVowel, arrVowel) * helper(countConsonant, arrConsonant);
    }

    public int helper(int m, int[] n) {
        long tmp = factorial(m);

        for (int num : n) {
            if (num != 0) {
                tmp /= factorial(num);
            }
        }

        return (int) tmp;
    }

    public long factorial(int n) {
        long answer = 1;
        for (int i = 1; i <= n; i++) {
            answer *= i;
        }

        return answer;
    }
}
