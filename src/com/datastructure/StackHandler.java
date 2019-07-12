/**
 * 
 */
package com.datastructure;

/**
 * @author Chris Lin
 * @version 1.0
 */
public abstract class StackHandler<T> {
    abstract public int size();
    abstract public boolean isEmpty();
    abstract public void push(T aValue);
    abstract public T pop();
    abstract public T peek();
    abstract public void printStack();
}
