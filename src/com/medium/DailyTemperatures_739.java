/**
 * 
 */
package com.medium;

import java.util.Stack;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class DailyTemperatures_739 {
    public static void main(String[] args) {
        DailyTemperatures_739 mTest = new DailyTemperatures_739();
        int[] T = new int[] { 73, 74, 75, 71, 69, 72, 76, 73 };
        int[] mResult = mTest.dailyTemperatures(T);
        for (int element : mResult) {
            System.out.print(element + ", ");
        }
    }

    // Use stack to record in reverse order
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 59 ms, faster than 32.32%
    // Memory Usage: 44.3 MB, less than 21.71%
    public int[] dailyTemperatures(int[] T) {
        if (T == null || T.length == 0) {
            return T;
        }

        int mLength = T.length;
        int[] mResult = new int[mLength];
        Stack<Integer> mStack = new Stack<>();
        for (int i = mLength - 1; i >= 0; i--) {
            while (!mStack.isEmpty() && T[i] >= T[mStack.peek()]) {
                // System.out.printf("!!!!!!Round:%d\n",i);
                mStack.pop();
            }

            mResult[i] = mStack.isEmpty() ? 0 : (mStack.peek() - i);
            // System.out.printf("peek:%d i: %d\n", mStack.isEmpty() ? 0
            // :mStack.peek(), i);

            mStack.push(i);
            // System.out.printf("push: %d\n", i);
        }

        return mResult;
    }
}
