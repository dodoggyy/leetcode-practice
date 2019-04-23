/**
 * 
 */
package com.easy;

/**
 * @author Chris
 *
 */
public class IntersectionofTwoLinkedLists_160 {

    /**
     * @param args
     */
    public static void main(String[] args) {
        // TODO Auto-generated method stub

    }

    // use cut long linklist part and traversal check
    // Runtime: 1 ms, faster than 98.12%
    // Memory Usage: 36.6 MB, less than 66.51%
    public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
        if (headA == null || headB == null) {
            return null;
        }
        int mLengthA = getLength(headA);
        int mLengthB = getLength(headB);

        if (mLengthA > mLengthB) {
            for (int i = 0; i < (mLengthA - mLengthB); i++) {
                headA = headA.next;
            }
        } else if (mLengthA < mLengthB) {
            for (int i = 0; i < (mLengthB - mLengthA); i++) {
                headB = headB.next;
            }
        }

        while (headA != null && headB != null && headA != headB) {
            headA = headA.next;
            headB = headB.next;
        }

        return (headA != null && headB != null) ? headA : null;

    }

    // use cycle check (A traversal end then link to B start and final may intersection)
    // A = a + c, B = b + c , a+c+b = a+b+c
    // Runtime: 1 ms, faster than 98.12%
    // Memory Usage: 40.1 MB, less than 7.45%
    public ListNode getIntersectionNode2(ListNode headA, ListNode headB) {
        if (headA == null || headB == null) {
            return null;
        }
        ListNode mListNode1 = headA, mListNode2 = headB;
        while (mListNode1 != mListNode2) {
            mListNode1 = (mListNode1 == null) ? headB : mListNode1.next;
            mListNode2 = (mListNode2 == null) ? headA : mListNode2.next;
        }
        return mListNode1;
    }

    public int getLength(ListNode aHead) {
        int mLength = 0;
        while (aHead != null) {
            mLength++;
            aHead = aHead.next;
        }
        return mLength;
    }
}
