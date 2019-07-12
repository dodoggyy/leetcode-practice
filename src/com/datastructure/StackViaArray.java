/**
 * 
 */
package com.datastructure;

/**
 * @author Chris Lin
 * @version 1.0
 */

@SuppressWarnings("unchecked")
public class StackViaArray<T> extends StackHandler<T> {
    private int mStackCapacity;

    private T[] mStack;
    private int mTop;
    

    public StackViaArray(int aCapacity) {
        this.mTop = 0;
        this.mStackCapacity = aCapacity;
        this.mStack = (T[]) new Object[aCapacity];
    }

    @Override
    public int size() {
        return mTop;
    }

    @Override
    public boolean isEmpty() {
        return (mTop == 0);
    }

    @Override
    public void push(T aValue) {
        if (mTop == (mStack.length)) {
            System.out.println("extend Array capacity:" + mStackCapacity);
            T[] mNewArray = (T[]) new Object[mStack.length + mStackCapacity];
            System.arraycopy(mStack, 0, mNewArray, 0, mStack.length);
            mStack = mNewArray;
        }

        mStack[mTop++] = aValue;
    }

    @Override
    public T pop() {
        if (isEmpty()) {
            System.out.println("stack empty");
            return null;
        } else {
            System.out.println("Top:" + mTop);
            T mValue = mStack[--mTop];
            mStack[mTop] = null;
            return mValue;
        }
    }

    @Override
    public void printStack() {
        if (mTop != 0) {
            for (T t : mStack) {
                if (t != null)
                    System.out.print(t + " ");
            }
        }
        System.out.println();

        System.out.printf("Empty:%b\n", isEmpty());
    }
    
    @Override
    public T peek() {
        // TODO Auto-generated method stub
        return mStack[mTop - 1];
    }

    public static void main(String[] args) {
        // TODO Auto-generated method stub
        StackViaArray<Integer> mStackObject = new StackViaArray<>(3);
        System.out.printf("pop item : %d\n", mStackObject.pop());
        mStackObject.printStack();
        mStackObject.push(1);
        System.out.printf("peek item : %d\n", mStackObject.peek());
        mStackObject.printStack();
        mStackObject.push(2);
        System.out.printf("peek item : %d\n", mStackObject.peek());
        mStackObject.push(3);
        System.out.printf("peek item : %d\n", mStackObject.peek());
        for (int i = 4; i < 10; i++) {
            mStackObject.push(i);
        }
        // mStackObject.push(4);
        mStackObject.printStack();
        System.out.printf("pop item : %d\n", mStackObject.pop());
        System.out.printf("pop item : %d\n", mStackObject.pop());
        mStackObject.printStack();
        
        System.out.printf("peek item : %d\n", mStackObject.peek());
        
    }

}
