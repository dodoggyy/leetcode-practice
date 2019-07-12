/**
 * 
 */
package com.datastructure;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class QueueViaLinkedList<T> extends QueueHandler<T> {
    private class ListNode {
        T mValue;
        ListNode mNext;

        private ListNode(T t) {
            this.mValue = t;
        }
    }

    private ListNode mRear;
    private ListNode mFront;

    public QueueViaLinkedList(int aCapacity) {
        // TODO Auto-generated constructor stub
        this.mRear = null;
        this.mFront = null;
    }

    @Override
    public int size() {
        // TODO Auto-generated method stub
        int mCount = 0;
        ListNode mTmpNode = mRear.mNext;
        while (mTmpNode != null) {
            mCount++;
            mTmpNode = mTmpNode.mNext;
            if (mTmpNode == mRear) {
                mCount++;
                break;
            }
        }
        return mCount;
    }

    @Override
    public boolean isEmpty() {
        // TODO Auto-generated method stub
        return ((mRear == mFront) && (mRear == null) && (mFront == null));
    }

    @Override
    public void enqueue(T aValue) {
        // TODO Auto-generated method stub
        ListNode mTmpNode = new ListNode(aValue);
        if(mFront == null) {
            mFront = mTmpNode;
        } else {
            mRear.mNext = mTmpNode;
        }
        mRear = mTmpNode;
        mRear.mNext = mFront;
    }

    @Override
    public T dequeue() {
        // TODO Auto-generated method stub
        if(isEmpty()) { // empty case
            System.out.println("Queue empty");
            return null;
        } else { // 1 node case
            T mValue = null;
            if(mFront == mRear) {
                mValue = mFront.mValue;
                mFront = null;
                mRear = null;
            } else { // 2 or more nodes case
                ListNode mTmpNode = mFront;
                mValue = mTmpNode.mValue;
                mFront = mFront.mNext;
                mRear.mNext = mFront;
            }
            return mValue;
        }
        
    }

    @Override
    public T peek() {
        // TODO Auto-generated method stub
        return mRear.mValue;
    }

    @Override
    public void printQueue() {
        // TODO Auto-generated method stub
        ListNode mTmpNode = mFront;
        while(mTmpNode != mRear) {
            System.out.print(mTmpNode.mValue + " ");
            mTmpNode = mTmpNode.mNext;
        }
        System.out.println(mTmpNode.mValue);
    }

    /**
     * @param args
     */
    public static void main(String[] args) {
        // TODO Auto-generated method stub
        QueueViaArray<Integer> mQueue = new QueueViaArray<>(5);
        mQueue.enqueue(1);
        mQueue.printQueue();
        mQueue.enqueue(2);
        mQueue.printQueue();
        mQueue.enqueue(3);
        mQueue.printQueue();
        System.out.println("dequeue: " + mQueue.dequeue());
        System.out.println("dequeue: " + mQueue.dequeue());
        System.out.println("dequeue: " + mQueue.dequeue());
        System.out.println("dequeue: " + mQueue.dequeue());

        QueueViaArray<String> mQueue2 = new QueueViaArray<String>(5);
        mQueue2.enqueue("A");
        mQueue2.printQueue();
        mQueue2.enqueue("B");
        mQueue2.printQueue();
        mQueue2.enqueue("C");
        mQueue2.printQueue();
        mQueue2.enqueue("D");
        mQueue2.printQueue();
        mQueue2.enqueue("E");
        mQueue2.printQueue();
        System.out.println("dequeue: " + mQueue2.dequeue());
        System.out.println("dequeue: " + mQueue2.dequeue());
        System.out.println("dequeue: " + mQueue2.dequeue());
        System.out.println("dequeue: " + mQueue2.dequeue());
        mQueue2.enqueue("F");
        mQueue2.printQueue();
        mQueue2.enqueue("G");
        mQueue2.printQueue();
        mQueue2.enqueue("H");
        mQueue2.printQueue();
        mQueue2.enqueue("I");
        mQueue2.printQueue();
    }

}
