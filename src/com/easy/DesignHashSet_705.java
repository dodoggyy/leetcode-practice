package com.easy;

import java.util.ArrayList;

public class DesignHashSet_705 {
    // 2019年8月24日
    // Use LinkedList concept
    // Time Complexity: O(1)
    // Space Complexity:O(n)
    // Runtime: 68 ms, faster than 20.33%
    // Memory Usage: 61.5 MB, less than 22.22%

    private static class LinkedListNode {
        public LinkedListNode next;
        public LinkedListNode pre;
        public int value;

        public LinkedListNode(int value) {
            this.value = value;
        }
    }

    private ArrayList<LinkedListNode> arr;

    class MyHashSet {
        private int capacity = 100;

        /** Initialize your data structure here. */
        public MyHashSet() {
            arr = new ArrayList<>();
            for (int i = 0; i < capacity; i++) {
                arr.add(null);
            }
        }

        private int getIndexForKey(Integer key) {
            return Math.abs(key.hashCode() % arr.size());
        }

        private LinkedListNode getNodeForKey(Integer key) {
            int mIndex = getIndexForKey(key);
            LinkedListNode mCurrent = arr.get(mIndex);
            while (mCurrent != null) {
                if (mCurrent.value == key) {
                    return mCurrent;
                }
                mCurrent = mCurrent.next;
            }
            return null;
        }

        public void add(int key) {
            LinkedListNode mNode = getNodeForKey(key);
            if (mNode == null) {
                mNode = new LinkedListNode(key);
                int mIndex = getIndexForKey(key);
                if (arr.get(mIndex) != null) {
                    mNode.next = arr.get(mIndex);
                    mNode.next.pre = mNode;
                }
                arr.set(mIndex, mNode);
            }
        }

        public void remove(int key) {
            LinkedListNode mNode = getNodeForKey(key);
            if (mNode != null) {
                if (mNode.pre != null) {
                    mNode.pre.next = mNode.next;
                } else {
                    int mHashKey = getIndexForKey(key);
                    arr.set(mHashKey, mNode.next);
                }

                if (mNode.next != null) {
                    mNode.next.pre = mNode.pre;
                }

            }
        }

        /** Returns true if this set contains the specified element */
        public boolean contains(int key) {
            return getNodeForKey(key) == null ? false : true;
        }
    }
}

/**
 * Your MyHashSet object will be instantiated and called as such: MyHashSet obj
 * = new MyHashSet(); obj.add(key); obj.remove(key); boolean param_3 =
 * obj.contains(key);
 */
