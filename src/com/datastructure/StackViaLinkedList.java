/**
 * 
 */
package com.datastructure;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class StackViaLinkedList<T> extends StackHandler<T> {

    /*
     * Stack linked list structure
     */
    private class ListNode {
        T mValue;
        ListNode mNext;

        public ListNode(T t) {
            this.mValue = t;
        }
    }

    private int mSize;
    private ListNode mHead;

    public StackViaLinkedList() {
        this.mSize = 0;
        mHead = null;
    }

    @Override
    public int size() {
        // TODO Auto-generated method stub
        return mSize;
    }

    @Override
    public boolean isEmpty() {
        // TODO Auto-generated method stub
        return (mSize == 0);
    }

    @Override
    public void push(T aValue) {
        // TODO Auto-generated method stub
        ListNode mNewNode = new ListNode(aValue);
        if (mHead == null) {
            mHead = mNewNode;
        } else {
            mNewNode.mNext = mHead;
            mHead = mNewNode;
        }
        mSize++;
    }

    @Override
    public T pop() {
        // TODO Auto-generated method stub
        if (mHead == null) {
            System.out.println("Stack empty");
            return null;
        } else {
            ListNode mTmpNode = mHead;
            mHead = mHead.mNext;
            mSize--;
            return mTmpNode.mValue;
        }
    }

    public T peek() {
        if (mHead == null) {
            return null;
        } else {
            return mHead.mValue;
        }
    }

    @Override
    public void printStack() {
        // TODO Auto-generated method stub
        if(isEmpty()) {
            System.out.println("Stack empty");
        } else {
            ListNode mTmp = mHead;
            while(mTmp != null) {
                System.out.print( mTmp.mValue +" ");
                mTmp = mTmp.mNext;
            }
        }
    }

    public static void main(String[] args) {
        // TODO Auto-generated method stub
        StackViaArray<Integer> mStackObject = new StackViaArray<>(3);
        System.out.printf("pop item : %d\n", mStackObject.pop());
        mStackObject.printStack();
        mStackObject.push(1);
        mStackObject.printStack();
        mStackObject.push(2);
        mStackObject.push(3);
        for (int i = 4; i < 10; i++) {
            mStackObject.push(i);
        }
        // mStackObject.push(4);
        mStackObject.printStack();
        System.out.printf("pop item : %d\n", mStackObject.pop());
        System.out.printf("pop item : %d\n", mStackObject.pop());
        mStackObject.printStack();
    }
}
