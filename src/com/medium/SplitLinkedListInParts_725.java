/**
 * 
 */
package com.medium;

/**
 * @author linquanliang
 *
 */
public class SplitLinkedListInParts_725 {

    // Split Input List:
    // Time Complexity: O(N+k)
    // Space Complexity:O(k)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 39.6 MB, less than 7.69%
    public ListNode[] splitListToParts(ListNode root, int k) {
        ListNode[] result = new ListNode[k];

        int size = 0;
        ListNode cur = root;
        while (cur != null) {
            cur = cur.next;
            size++;
        }
        int n = size / k, rem = size % k;

        cur = root;
        ListNode pre = null;
        for (int i = 0; i < k; i++, rem--) {
            if (cur == null) {
                break;
            }
            result[i] = cur;
            for (int j = 0; j < n + (rem > 0 ? 1 : 0); j++) {
                pre = cur;
                cur = cur.next;
            }
            pre.next = null;
        }

        return result;
    }

}
