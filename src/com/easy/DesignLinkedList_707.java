/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class DesignLinkedList_707 {
    // 2019年8月29日
    // Use own double LinkedList
    // Time Complexity: O(1)
    // Space Complexity:O(n)
    // Runtime: 49 ms, faster than 99.05%
    // Memory Usage: 45.2 MB, less than 88.89%
    public static class ListNode {
        int val;
        ListNode next;
        ListNode pre;

        public ListNode(int value) {
            this.val = value;
            this.next = null;
            this.pre = null;
        }
    }

    private ListNode head;
    private ListNode tail;
    private int size = 0;

    /** Initialize your data structure here. */
    public DesignLinkedList_707() {
        head = null;
        tail = null;
    }

    /**
     * Get the value of the index-th node in the linked list. If the index is
     * invalid, return -1.
     */
    // O(n)
    public int get(int index) {
        if (index < 0 || index >= size) {
            return -1;
        }
        ListNode mTmpNode = head;
        for (int i = 0; i < index; i++) {
            mTmpNode = mTmpNode.next;
        }
        return mTmpNode.val;
    }

    /**
     * Add a node of value val before the first element of the linked list.
     * After the insertion, the new node will be the first node of the linked
     * list.
     */
    // O(1)
    public void addAtHead(int val) {
        ListNode mTmpNode = new ListNode(val);
        if (size == 0) {
            head = mTmpNode;
            tail = mTmpNode;
        }
        head.pre = mTmpNode;
        mTmpNode.next = head;
        head = mTmpNode;
        size++;
    }

    /** Append a node of value val to the last element of the linked list. */
    // O(1)
    public void addAtTail(int val) {
        ListNode mTmpNode = new ListNode(val);
        if (size == 0) {
            head = mTmpNode;
            tail = mTmpNode;
        }
        tail.next = mTmpNode;
        mTmpNode.pre = tail;
        tail = mTmpNode;
        size++;
    }

    /**
     * Add a node of value val before the index-th node in the linked list. If
     * index equals to the length of linked list, the node will be appended to
     * the end of linked list. If index is greater than the length, the node
     * will not be inserted.
     */
    // O(n)
    public void addAtIndex(int index, int val) {
        if (index > size) {
            return;
        }

        if (index <= 0) {
            addAtHead(val);
        } else if (index == size) {
            addAtTail(val);
        } else {
            ListNode mTmpNode = head;
            for (int i = 0; i < index; i++) {
                mTmpNode = mTmpNode.next;
            }
            ListNode mNode = new ListNode(val);
            mTmpNode.pre.next = mNode;
            mNode.pre = mTmpNode.pre;
            mTmpNode.pre = mNode;
            mNode.next = mTmpNode;
            size++;
        }
    }

    /** Delete the index-th node in the linked list, if the index is valid. */
    // O(1)
    public void deleteAtIndex(int index) {
        if (index < 0 || index >= size) {
            return;
        }
        if (index == 0) {
            if (size > 0) {
                head = head.next;
                size--;
            }
            if (size == 0) {
                tail = null;
            }
        } else {
            ListNode mTmpNode = head;
            for (int i = 0; i < index; i++) {
                mTmpNode = mTmpNode.next;
            }
            if (mTmpNode == tail) {
                mTmpNode.pre.next = null;
                tail = tail.pre;
            } else {
                mTmpNode.pre.next = mTmpNode.next;
                mTmpNode.next.pre = mTmpNode.pre;
            }
            size--;
        }

    }
}
