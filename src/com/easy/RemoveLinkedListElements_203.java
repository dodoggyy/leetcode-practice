/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class RemoveLinkedListElements_203 {
    // 2019年7月11日
    // Use tmp ListNode traversal
    // Time Complexity: O()
    // Space Complexity:O()
    // Runtime: 1 ms, faster than 98.64%
    // Memory Usage: 39.2 MB, less than 99.95%
    public ListNode removeElements(ListNode head, int val) {
        if (head == null) {
            return null;
        }
        while (head != null && head.val == val) {
            head = head.next;
        }
        ListNode mNode = head;
        while (mNode != null && mNode.next != null) {
            if (mNode.next.val == val) {
                mNode.next = mNode.next.next;
            } else {
                mNode = mNode.next;
            }
        }
        return head;
    }
}
