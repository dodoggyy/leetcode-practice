/**
 * 
 */
package com.medium;

import java.util.HashMap;
import java.util.LinkedHashMap;
import java.util.Map;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class LRUCache_146 {
    // 2019年8月6日
    // Use Linked HashMap
    // Time Complexity: O(1)
    // Space Complexity:O(n)
    // Runtime: 61 ms, faster than 68.53%
    // Memory Usage: 57.3 MB, less than 67.17%
    private LinkedHashMap<Integer, Integer> mMap;
    private int capacity;

    public LRUCache_146(int capacity) {
        this.capacity = capacity;
        mMap = new LinkedHashMap<Integer, Integer>(capacity, 0.75f, true) {
            protected boolean removeEldestEntry(Map.Entry eldest) {
                return size() > capacity;
            }
        };
    }

    public int get(int key) {
        return mMap.getOrDefault(key, -1);
    }

    public void put(int key, int value) {
        mMap.put(key, value);
    }
}

// 2019年8月6日
// Use HashMap and double linkedlist
// Time Complexity: O(1)
// Space Complexity:O(n)
// Runtime: 62 ms, faster than 58.49%
// Memory Usage: 57.9 MB, less than 56.99%
class LRUCache {
    class Node {
        int key;
        int value;
        Node pre;
        Node next;

        public Node(int key, int value) {
            this.key = key;
            this.value = value;
        }
    }

    private Node head;
    private Node tail;
    private int capacity;
    private HashMap<Integer, Node> mMap;

    public LRUCache(int capacity) {
        this.capacity = capacity;
        mMap = new HashMap<>();
        head = new Node(0, 0);
        tail = new Node(0, 0);
        head.next = tail;
        tail.pre = head;
        head.pre = null;
        tail.next = null;
    }

    public void addNode2Head(Node mNode) {
        Node mNodeHeadNext = head.next;
        head.next = mNode;
        mNodeHeadNext.pre = mNode;
        mNode.next = mNodeHeadNext;
        mNode.pre = head;
    }

    public void deleteNode(Node mNode) {
        mNode.pre.next = mNode.next;
        mNode.next.pre = mNode.pre;
    }

    public int get(int key) {
        if (mMap.containsKey(key)) {
            Node mNode = mMap.get(key);
            int mResult = mNode.value;
            deleteNode(mNode);
            addNode2Head(mNode);

            return mResult;
        }
        return -1;
    }

    public void put(int key, int value) {
        if (mMap.containsKey(key)) {
            Node mNode = mMap.get(key);
            deleteNode(mNode);
            mNode.value = value;
            addNode2Head(mNode);
        } else {
            if (mMap.size() == capacity) {
                mMap.remove(tail.pre.key);
                deleteNode(tail.pre);
            }
            Node mNode = new Node(key, value);
            addNode2Head(mNode);
            mMap.put(key, mNode);

        }
    }

}
