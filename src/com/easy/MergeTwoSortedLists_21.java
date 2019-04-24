/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class MergeTwoSortedLists_21 {

    // Iterative version
    // Runtime: 1 ms, faster than 86.30%
    // Memory Usage: 37.1 MB, less than 97.90%
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        if (l1 == null) {
            return l2;
        }
        if (l2 == null) {
            return l1;
        }
        ListNode mNodeMerge = new ListNode(0);
        ListNode mNodeMergeP = mNodeMerge;
        ListNode mNodeTmp1 = l1;
        ListNode mNodeTmp2 = l2;
        while (mNodeTmp1 != null && mNodeTmp2 != null) {
            if (mNodeTmp1.val < mNodeTmp2.val) {
                mNodeMergeP.next = mNodeTmp1;
                mNodeTmp1 = mNodeTmp1.next;
            } else {
                mNodeMergeP.next = mNodeTmp2;
                mNodeTmp2 = mNodeTmp2.next;
            }
            mNodeMergeP = mNodeMergeP.next;
        }
        if (mNodeTmp1 != null) {
            mNodeMergeP.next = mNodeTmp1;
        }
        if (mNodeTmp2 != null) {
            mNodeMergeP.next = mNodeTmp2;
        }

        return mNodeMerge.next;
    }
    
    // Recursive version
    // Runtime: 0 ms, faster than 100.00% 
    // Memory Usage: 37.2 MB, less than 97.90%
    public ListNode mergeTwoLists2(ListNode l1, ListNode l2) {
        if (l1 == null) {
            return l2;
        }
        if (l2 == null) {
            return l1;
        }
        
        if(l1.val < l2.val) {
            l1.next = mergeTwoLists2(l1.next, l2);
            return l1;
        } else {
            l2.next = mergeTwoLists2(l1, l2.next);
            return l2;
        }
    }
}
