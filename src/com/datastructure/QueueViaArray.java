/**
 * 
 */
package com.datastructure;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class QueueViaArray<T> extends QueueHandler<T> {
    private int mQueueCapacity;

    private int mRear;
    private int mFront;
    private boolean bIsQueueFull;
    private T[] mQueue;

    @SuppressWarnings("unchecked")
    public QueueViaArray(int aQueueCapacity) {
        this.mQueueCapacity = aQueueCapacity;
        mQueue = (T[]) new Object[mQueueCapacity];
        this.mRear = -1;
        this.mFront = -1;
        this.bIsQueueFull = false;
    }

    @Override
    public int size() { // O(1)
        // TODO Auto-generated method stub
        return (mRear - mFront + mQueueCapacity) % mQueueCapacity;
    }

    @Override
    public boolean isEmpty() {
        // TODO Auto-generated method stub
        return (!bIsQueueFull && (mRear == mFront));
    }

    @Override
    public void enqueue(T aValue) {
        // TODO Auto-generated method stub
        if (bIsQueueFull) {
            System.out.println("Queue full! cannot enqueue");
        } else {
            mRear = (mRear + 1) % mQueueCapacity;
            mQueue[mRear] = aValue;
            if (mRear == mFront) {
                bIsQueueFull = true;
            }
            if(mFront == -1) {
                mFront++;
            }
        }
    }

    @Override
    public T dequeue() {
        // TODO Auto-generated method stub
        if (isEmpty()) {
            System.out.println("Queue empty! cannot dequeue");
            return null;
        } else {
            mFront = (mFront + 1) % mQueueCapacity;
            T mElement = mQueue[mFront];
            if (mFront == mRear) {
                bIsQueueFull = false;
            }
            return mElement;
        }

    }

    @Override
    public T peek() {
        // TODO Auto-generated method stub
        return mQueue[mRear];
    }

    @Override
    public void printQueue() {
        // TODO Auto-generated method stub
        if (isEmpty()) {
            System.out.println("queue empty");
        } else {
            int mIndex = (mFront + 1)% mQueueCapacity;
            while (mIndex != mRear) {
                System.out.print(mQueue[mIndex] + " ");
                mIndex = (mIndex + 1) % mQueueCapacity;
            }
            System.out.print(mQueue[mIndex] + " ");
            System.out.println();
        }

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
