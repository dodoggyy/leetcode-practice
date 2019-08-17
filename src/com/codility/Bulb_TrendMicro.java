/**
 * 
 */
package com.codility;

import java.util.HashSet;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class Bulb_TrendMicro {

    /**
     * @param args
     */
    public static void main(String[] args) {
        // TODO Auto-generated method stub
        Bulb_TrendMicro mTest = new Bulb_TrendMicro();
        int[] A = new int[] { 2, 1, 3, 5, 4 };
        int[] A2 = new int[] { 1, 2, 3, 5, 4 };
        System.out.println(mTest.solution(A));
        System.out.println(mTest.solution2(A));

    }

    // 2019年8月18日
    // Use HashSet to memorize
    // Time Complexity: O(n^2)
    // Space Complexity:O(n)
    public int solution(int[] A) {
        // write your code in Java SE 8
        int mResult = 0;
        HashSet<Integer> mSet = new HashSet<>();

        for (int num : A) {
            boolean bIsShine = true;
            for (int i = 1; i < num; i++) {
                if (!mSet.contains(i)) {
                    bIsShine = false;
                    break;
                }
            }
            if (bIsShine) {
                mResult++;
            }
            mSet.add(num);
        }

        return mResult;
    }

    // 2019年8月18日
    // Use sum check
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    public int solution2(int[] A) {
        // write your code in Java SE 8
        int mResult = 0;
        int[] mTmpSum = new int[A.length];
        for (int num : A) {

            for (int i = num; i < A.length; i++) {
                mTmpSum[i - 1] += num;
            }
            int mTmpTotal = ((num + 1) * num) / 2;
//            System.out.println(num + ":" + mTmpTotal + ":" + mTmpSum[num - 1]);
            if (mTmpSum[num - 1] == mTmpTotal) {
                mResult++;
            }
        }
        return mResult;
    }

}
