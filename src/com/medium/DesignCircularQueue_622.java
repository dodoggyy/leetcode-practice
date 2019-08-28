/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class DesignCircularQueue_622 {
    // 2019年8月29日
    // Use array
    // Time Complexity: O(1)
    // Space Complexity:O(n)
    // Runtime: 44 ms, faster than 100.00%
    // Memory Usage: 38 MB, less than 28.57%

    private int rear;
    private int front;
    private int size;
    private int capacity;
    private int[] mQueue;

    /**
     * Initialize your data structure here. Set the size of the queue to be k.
     */
    public DesignCircularQueue_622(int k) {
        this.capacity = k;
        this.size = 0;
        this.rear = -1;
        this.front = -1;
        mQueue = new int[capacity];
    }

    /**
     * Insert an element into the circular queue. Return true if the operation
     * is successful.
     */
    public boolean enQueue(int value) {
        if (isFull()) {
            return false;
        }
        if(front == -1) {
            front++;
        }
        rear = (rear + 1) % capacity;
        mQueue[rear] = value;
        size++;
        return true;
    }

    /**
     * Delete an element from the circular queue. Return true if the operation
     * is successful.
     */
    public boolean deQueue() {
        if (isEmpty()) {
            return false;
        }
        front = (front + 1) % capacity;
        size--;
        return true;
    }

    /** Get the front item from the queue. */
    public int Front() {
        if (!isEmpty()) {
            return mQueue[front];
        }
        return -1;
    }

    /** Get the last item from the queue. */
    public int Rear() {
        if (!isEmpty()) {
            return mQueue[rear];
        }
        return -1;
    }

    /** Checks whether the circular queue is empty or not. */
    public boolean isEmpty() {
        return size == 0;
    }

    /** Checks whether the circular queue is full or not. */
    public boolean isFull() {
        return size == capacity;
    }
}
