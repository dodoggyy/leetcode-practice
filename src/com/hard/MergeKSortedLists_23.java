/**
 * 
 */
package com.hard;

import java.util.Comparator;
import java.util.PriorityQueue;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class MergeKSortedLists_23 {
    // 2019年9月8日
    // Use PriorityQueue
    // Time Complexity: O(nlogk)
    // Space Complexity:O(n)
    // n for total node, k for number of linked list
    // Runtime: 6 ms, faster than 55.54%
    // Memory Usage: 43.4 MB, less than 26.78%
    public ListNode mergeKLists(ListNode[] lists) {
        if (lists == null || lists.length == 0) {
            return null;
        }

        PriorityQueue<ListNode> mQueue = new PriorityQueue<>(lists.length, new Comparator<ListNode>() {

            @Override
            public int compare(ListNode o1, ListNode o2) {
                // TODO Auto-generated method stub
                return o1.val - o2.val;
            }
        });

        for (ListNode node : lists) {
            if (node != null) {
                mQueue.add(node);
            }
        }

        ListNode head = new ListNode(0);
        ListNode mPoint = head;
        while (!mQueue.isEmpty()) {
            mPoint.next = mQueue.poll();
            mPoint = mPoint.next;
            ListNode mNext = mPoint.next;
            if (mNext != null) {
                mQueue.add(mNext);
            }
        }

        return head.next;
    }

    // 2019年9月8日
    // Use divide and conquer
    // Time Complexity: O(nlogk)
    // Space Complexity:O(1)
    // n for total node, k for number of linked list
    // Runtime: 3 ms, faster than 84.54%
    // Memory Usage: 42.4 MB, less than 27.87%
    public ListNode mergeKLists2(ListNode[] lists) {
        return partition(lists, 0, lists.length - 1);
    }

    private ListNode partition(ListNode[] lists, int mStart, int mEnd) {
        if (mStart == mEnd) {
            return lists[mStart];
        }
        if (mStart < mEnd) {
            int mMid = mStart + (mEnd - mStart) / 2;
            ListNode l1 = partition(lists, mStart, mMid);
            ListNode l2 = partition(lists, mMid + 1, mEnd);
            return merge(l1,l2);
        } else {
            return null;
        }
    }
    
    private ListNode merge(ListNode l1, ListNode l2) {
        if(l1 == null) {
            return l2;
        }
        if(l2 == null) {
            return l1;
        }
        if(l1.val < l2.val) {
            l1.next = merge(l1.next, l2);
            return l1;
        } else {
            l2.next = merge(l1, l2.next);
            return l2;
        }
    }
}
