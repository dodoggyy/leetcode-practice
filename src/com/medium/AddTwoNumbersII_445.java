/**
 * 
 */
package com.medium;

import java.util.Stack;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class AddTwoNumbersII_445 {
    // 2019年12月11日
    // Use stack
    // Time Complexity: O(MAX(m,n))
    // Space Complexity:O(MAX(m,n))
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 0.0 MB, less than 100%
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode mResult = new ListNode(0);
        ListNode head = mResult;
        Stack<Integer> mStack1 = new Stack<>();
        Stack<Integer> mStack2 = new Stack<>();

        while (l1 != null) {
            mStack1.push(l1.val);
            l1 = l1.next;
        }

        while (l2 != null) {
            mStack2.push(l2.val);
            l2 = l2.next;
        }

        int mCarry = 0;
        while (!mStack1.isEmpty() || !mStack2.isEmpty() || mCarry != 0) {

            int num1 = mStack1.isEmpty() ? 0 : mStack1.pop();
            int num2 = mStack2.isEmpty() ? 0 : mStack2.pop();
            int mSum = num1 + num2 + mCarry;
            ListNode mNode = new ListNode(mSum % 10);
            mNode.next = head.next;
            head.next = mNode;
            mCarry = mSum / 10;
        }

        return mResult.next;
    }
}
