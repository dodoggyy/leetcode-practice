/**
 * 
 */
package com.codility;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class NthPowerOf11_TrendMicro {

    /**
     * @param args
     */
    public static void main(String[] args) {
        // TODO Auto-generated method stub
        NthPowerOf11_TrendMicro mTest = new NthPowerOf11_TrendMicro();
        int N = 66;
        System.out.println(mTest.solution(N));
        System.out.println(mTest.solution2(N));
    }

    // 2019年8月17日
    // Use brute force
    // Time Complexity: O(11^n)
    // Space Complexity:O(1)
    // only can handle 11^7
    public int solution(int N) {
        // write your code in Java SE 8
        int mResult = 0;

        long mNumber = (long) Math.pow(11, N);
        String mStr = mNumber + "";

        for (char c : mStr.toCharArray()) {
            if (c == '1') {
                mResult++;
            }
        }

        return mResult;
    }

    // 2019年8月17日
    // Use Pascal's triangle
    // Time Complexity: O(log(11^n))
    // Space Complexity:O(log(11^n))
    // Ref:https://bitbucket.org/dodoggyy/leetcode/src/master/src/com/easy/PascalsTriangleII_119.java
    // only can handle 11^66
    public int solution2(int N) {
        // write your code in Java SE 8
        int mResult = 0;
        long[] mDigits = new long[N + 1];

        mDigits[0] = 1;
        for (int i = 1; i <= N; i++) {
            for (int j = i; j > 0; j--) {
                mDigits[j] += mDigits[j - 1];
            }
        }

        long mCarry = 0;
        for (int i = mDigits.length - 1; i > 0; i--) {
            mDigits[i] += mCarry;
            mCarry = mDigits[i] / 10;
            mDigits[i] %= 10;
        }

        for (long i : mDigits) {
            // System.out.print(i + " ");
            if (i == 1) {
                mResult++;
            }
        }

        while (mCarry != 0) {
            if (mCarry % 10 == 0) {
                mResult++;
            }
            mCarry /= 10;
        }

        return mResult;
    }

    // 2019年8月17日
    // Use array shift
    // Time Complexity: O()
    // Space Complexity:O()
    public int solution3(int N) {
        int mResult = 0;
        List<Character> mDigit = new ArrayList<>();
        mDigit.add('1');
        mDigit.add('1');

        for (int i = 1; i < N; i++) {
//            List<Character> mOrigDigit = ((Object) mDigit).clone();
//            mDigit[]
        }

        return mResult;
    }
}
