/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ReverseLinkedListII_92 {
    // 2020年3月14日
    // Use iterative:
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 37.3 MB, less than 11.36%
    public ListNode reverseBetween(ListNode head, int m, int n) {
        ListNode dummy = new ListNode(-1);
        dummy.next = head;
        ListNode pre = dummy;
        ListNode cur = null;
        for (int i = 0; i < m - 1; i++) {
            pre = pre.next;
        }
        cur = pre.next;
        for (int i = m; i < n; i++) {
            ListNode tmp = cur.next;
            cur.next = tmp.next;
            tmp.next = pre.next;
            pre.next = tmp;
        }
        
        return dummy.next;
    }
}
