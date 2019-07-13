/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 * 
 */
public class DeleteNodeInALinkedList_237 {
    // Copy next node to current and link to next next node
    // Time Complexity: O(1)
    // Space Complexity: O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 36.8 MB, less than 99.82% 
    public void deleteNode(ListNode node) {
        node.val = node.next.val;
        node.next = node.next.next;
    }
}
