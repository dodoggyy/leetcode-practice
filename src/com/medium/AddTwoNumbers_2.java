/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class AddTwoNumbers_2 {
    public static void main(String[] args) throws Exception {
        ListNode l1 = new ListNode(1);
        ListNode l2 = new ListNode(9);
        l2.next = new ListNode(9);

        AddTwoNumbers_2 mCalculator = new AddTwoNumbers_2();

        ListNode mResult = mCalculator.addTwoNumbers(l1, l2);
        while (mResult != null) {
            System.out.print(mResult.val + "->");
            mResult = mResult.next;
        }
    }

    // Use general math calculation
    // Need judge carry bit
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 2 ms, faster than 87.24%
    // Memory Usage: 44.8 MB, less than 84.40%
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode mResult = new ListNode(-1);
        ListNode mNodeIndex = mResult;
        int mTmp = 0;
        while (l1 != null || l2 != null || mTmp != 0) {
            mTmp += ((l1 == null) ? 0 : l1.val) + ((l2 == null) ? 0 : l2.val);
            ListNode mNode = new ListNode(mTmp % 10);
            mNodeIndex.next = mNode;
            mNodeIndex = mNodeIndex.next;
            mTmp /= 10;
            if (l1 != null) {
                l1 = l1.next;
            }
            if (l2 != null) {
                l2 = l2.next;
            }
        }
        return mResult.next;
    }
}
