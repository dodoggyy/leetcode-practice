package com.datastructure;

import java.util.ArrayList;
import java.util.LinkedList;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class MyHashTable<K, V> {

    private static class LinkedListNode<K, V> {
        public LinkedListNode<K, V> next;
        public LinkedListNode<K, V> pre;
        public K key;
        public V value;

        public LinkedListNode(K k, V v) {
            this.key = k;
            this.value = v;
        }
    }

    private ArrayList<LinkedListNode<K, V>> arr;

    public MyHashTable(int mCapacity) {
        arr = new ArrayList<LinkedListNode<K, V>>();
        for (int i = 0; i < mCapacity; i++) {
            arr.add(null);
        }
    }

    public V put(K key, V value) {
        LinkedListNode<K, V> mNode = getNodeForKey(key);
        if (mNode != null) {
            V oldValue = mNode.value;
            mNode.value = value;
            return oldValue;
        }

        mNode = new LinkedListNode<K, V>(key, value);
        int mIndex = getIndexForKey(key);
        if (arr.get(mIndex) != null) {
            mNode.next = arr.get(mIndex);
            mNode.next.pre = mNode;
        }
        arr.set(mIndex, mNode);
        return null;
    }

    private int getIndexForKey(K key) {
        // TODO Auto-generated method stub
        return Math.abs(key.hashCode() % arr.size());
    }

    private LinkedListNode<K, V> getNodeForKey(K key) {
        // TODO Auto-generated method stub
        int mIndex = getIndexForKey(key);
        LinkedListNode<K, V> mCurrent = arr.get(mIndex);
        while (mCurrent != null) {
            if (mCurrent.key == key) {
                return mCurrent;
            }
            mCurrent = mCurrent.next;
        }
        return null;
    }

    public V remove(K key) {
        LinkedListNode<K, V> mNode = getNodeForKey(key);
        if (mNode != null) {
            if (mNode.pre != null) {
                mNode.pre.next = mNode.next;
            } else {
                // delete head & update
                int mHashKey = getIndexForKey(key);
                arr.set(mHashKey, mNode.next);
            }

            if (mNode.next != null) {
                mNode.next.pre = mNode.pre;
            }

            return mNode.value;

        } else {
            return null;
        }

    }

    public V get(K key) {
        if (key == null) {
            return null;
        }
        LinkedListNode<K, V> mNode = getNodeForKey(key);

        return mNode == null ? null : mNode.value;
    }

    public static void main(String[] args) {
        int mCapacity = 16;
        MyHashTable<Integer, Integer> mHashTable = new MyHashTable<>(mCapacity);

        for (int i = 0; i < 10; i++) {
            int mValue = i * 10;
            System.out.println("Put key:" + i + " value:" + mValue);
            mHashTable.put(i, mValue);
        }

        System.out.println();

        for (int i = 0; i < 11; i++) {
            System.out.println("Get key:" + i + " value:" + mHashTable.get(i));
        }

        int mUpdateKey = 3, mUpateValue = 1000;
        System.out.printf("Upadte key: %d Origin Value: %d Update Value:%d\n", mUpdateKey,
                mHashTable.put(mUpdateKey, mUpateValue), mHashTable.get(mUpdateKey));

        int mRemoveKey = 4;
        System.out.printf("Remove key:%d Before Value:%d After Value:%d\n", mRemoveKey, mHashTable.remove(mRemoveKey),
                mHashTable.get(mRemoveKey));
    }
}
