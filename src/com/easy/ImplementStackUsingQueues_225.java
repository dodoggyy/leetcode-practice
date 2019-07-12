/**
 * 
 */
package com.easy;

import java.util.LinkedList;
import java.util.Queue;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ImplementStackUsingQueues_225 {

    // Use utility lib queue:
    // Time Complexity: NA
    // Space Complexity: NA
    // Runtime: 42 ms, faster than 96.82%
    // Memory Usage: 33 MB, less than 95.26%

    /**
     * Your MyStack object will be instantiated and called as such: MyStack obj
     * = new MyStack(); obj.push(x); int param_2 = obj.pop(); int param_3 =
     * obj.top(); boolean param_4 = obj.empty();
     */

    class MyStack {
        private Queue<Integer> mQueue;

        /** Initialize your data structure here. */
        public MyStack() {
            mQueue = new LinkedList<>();
        }

        /** Push element x onto stack. */
        public void push(int x) {
            mQueue.add(x);
            int mCount = mQueue.size();
            while (mCount > 1) {
                mQueue.add(mQueue.poll());
                mCount--;
            }
        }

        /** Removes the element on top of the stack and returns that element. */
        public int pop() {
            return mQueue.poll();
        }

        /** Get the top element. */
        public int top() {
            return mQueue.peek();
        }

        /** Returns whether the stack is empty. */
        public boolean empty() {
            return mQueue.isEmpty();
        }
    }
}
