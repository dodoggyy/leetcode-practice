/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 * 
 */
public class PalindromeLinkedList_234 {
    // Use 2 pointers to find the middle node, then head compare with reverse slow
    // Time Complexity: O(n)
    // Space Complexity: O(1)
    // Runtime: 1 ms, faster than 99.21%
    // Memory Usage: 39.8 MB, less than 99.90%
    public boolean isPalindrome(ListNode head) {
        ListNode mNodeFast = head;
        ListNode mNodeSlow = head;
        while (mNodeFast != null && mNodeFast.next != null) {
            mNodeFast = mNodeFast.next.next;
            mNodeSlow = mNodeSlow.next;
        }
        mNodeSlow = reverse(mNodeSlow);
        while (mNodeSlow != null) {
            if (mNodeSlow.val != head.val) {
                return false;
            }
            mNodeSlow = mNodeSlow.next;
            head = head.next;
        }

        return true;
    }

    private ListNode reverse(ListNode head) {
        ListNode mNodePre = null;
        ListNode mNodeNext = null;

        while (head != null) {
            mNodeNext = head.next;
            head.next = mNodePre;
            mNodePre = head;
            head = mNodeNext;
        }

        return mNodePre;
    }
}
