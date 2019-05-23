/**
 * 
 */
package com.datastructure;

/**
 * @author Chris
 *
 */
public abstract class StackHandler<T> {
    abstract public int size();
    abstract public boolean isEmpty();
    abstract public void push(T aValue);
    abstract public T pop();
    abstract public void printStack();
}
