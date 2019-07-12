/**
 * 
 */
package com.easy;

import java.util.Stack;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class MinStack_155 {
    // Use additional stack to store minimum value
    // Time Complexity: NA
    // Space Complexity:NA
    // Runtime: 50 ms, faster than 18.45%
    // Memory Usage: 40.1 MB, less than 32.43%
    public class MinStack {
        private Stack<Integer> mStack;
        private Stack<Integer> mMinStack; // store min value

        /** initialize your data structure here. */
        public MinStack() {
            mStack = new Stack<>();
            mMinStack = new Stack<>();
        }

        public void push(int x) {
            mStack.push(x);
            if (mMinStack.isEmpty() || mMinStack.peek() >= x) {
                mMinStack.push(x);
            }
        }

        public void pop() {
            int x = mStack.pop();
            if (mMinStack.peek() == x)
                mMinStack.pop();
        }

        public int top() {
            return mStack.peek();
        }

        public int getMin() {
            return mMinStack.peek();
        }
    }

    // Use only one stack and 1 parameter to store minimum value
    // Time Complexity: NA
    // Space Complexity:NA
    // Runtime: 45 ms, faster than 98.97%
    // Memory Usage: 40.7 MB, less than 30.99%
    public class MinStack2 {
        private Stack<Integer> mStack;
        private int mMin;

        /** initialize your data structure here. */
        public MinStack2() {
            mStack = new Stack<>();
            mMin = Integer.MAX_VALUE;
        }

        public void push(int x) {
            if (x <= mMin) {
                mStack.push(mMin);
                mMin = x;
            }
            mStack.push(x);
        }

        public void pop() {
            if (mStack.pop() == mMin) {
                mMin = mStack.pop();
            }
        }

        public int top() {
            return mStack.peek();
        }

        public int getMin() {
            return mMin;
        }
    }
}
