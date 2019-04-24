/**
 * 
 */
package com.medium;

import com.easy.ListNode;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class RemoveNthNodeFromEndofList_19 {
    
    // Two pass version
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 37 MB, less than 92.46%
    public ListNode removeNthFromEnd(ListNode head, int n) {
        if (head.next == null) {
            return null;
        }

        ListNode mNodeIndex = head;
        int mListLength = 0;
        while (mNodeIndex != null) {
            mNodeIndex = mNodeIndex.next;
            mListLength++;
        }
        if(mListLength == n) {
            return head.next;
        }
        mNodeIndex = head;
        for(int i = 0; i < (mListLength-n-1); i++) {
            mNodeIndex = mNodeIndex.next;
        }
        mNodeIndex.next = mNodeIndex.next.next;

        return head;
    }
 
    // One pass version
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 37 MB, less than 92.46%
    public ListNode removeNthFromEnd2(ListNode head, int n) {
        if (head.next == null) {
            return null;
        }
        
        ListNode mNodeFast = head;
        ListNode mNodeSlow = head;
        
        while(n > 0) {
            mNodeFast = mNodeFast.next;
            n--;
        }
        if(mNodeFast == null) {
            return head.next;
        }
        while(mNodeFast.next != null) {
            mNodeFast = mNodeFast.next;
            mNodeSlow = mNodeSlow.next;
        }
        mNodeSlow.next = mNodeSlow.next.next;
        return head;
    }
}
