/**
 * 
 */
package com.datastructure;

/**
 * @author Chris Lin
 * @version 1.0
 */
public abstract class QueueHandler<T> {
    abstract public int size();
    abstract public boolean isEmpty();
    abstract public void enqueue(T aValue);
    abstract public T dequeue();
    abstract public T peek();
    abstract public void printQueue();
}
