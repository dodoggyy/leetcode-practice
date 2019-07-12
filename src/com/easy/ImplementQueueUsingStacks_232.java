/**
 * 
 */
package com.easy;

import java.util.Stack;

/**
 * @author Chris Lin
 * @version 1.0
 */
@SuppressWarnings("unchecked")

public class ImplementQueueUsingStacks_232 {
    
    // Use utility lib stack:
    // Time Complexity: NA
    // Space Complexity: NA
    // Runtime: 41 ms, faster than 98.06%
    // Memory Usage: 33.4 MB, less than 57.77% 

    /**
     * Your MyQueue object will be instantiated and called as such: MyQueue obj
     * = new MyQueue(); obj.push(x); int param_2 = obj.pop(); int param_3 =
     * obj.peek(); boolean param_4 = obj.empty();
     */

    @SuppressWarnings("rawtypes")
    class MyQueue {
        private Stack mInStack;
        private Stack mOutStack;

        /** Initialize your data structure here. */
        public MyQueue() {
            mInStack = new Stack();
            mOutStack = new Stack();
        }

        /** Push element x to the back of queue. */
        public void push(int x) {
            mInStack.push(x);
        }

        /**
         * Removes the element from in front of queue and returns that element.
         */
        public int pop() {
            in2OutStack();
            return (int) mOutStack.pop();
        }

        /** Get the front element. */
        public int peek() {
            in2OutStack();
            return (int) mOutStack.peek();

        }

        private void in2OutStack() {
            if (mOutStack.isEmpty()) {
                while (!mInStack.isEmpty()) {
                    mOutStack.push(mInStack.pop());
                }
            }
        }

        /** Returns whether the queue is empty. */
        public boolean empty() {
            return (mInStack.isEmpty() && mOutStack.isEmpty());

        }
    }
}
