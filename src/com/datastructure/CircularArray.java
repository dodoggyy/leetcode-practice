/**
 * 
 */
package com.datastructure;

import java.util.Iterator;

/**
 * @author Chris Lin
 * @version 1.0
 * 
 */
public class CircularArray<T> implements Iterable<T> {
    private T[] items;
    private int mHead = 0;

    public CircularArray(int mSize) {
        this.items = (T[]) new Object[mSize];
    }

    private int convert(int mIndex) {
        if (mIndex < 0) {
            mIndex += items.length;
        }
        return (mHead + mIndex) % items.length;
    }

    public void rotate(int mShiftRight) {
        mHead = convert(mShiftRight);
    }

    public T get(int i) {
        if (i < 0 || i >= items.length) {
            throw new java.lang.IndexOutOfBoundsException();
        }
        return items[convert(i)];
    }

    public void set(int i, T item) {
        items[convert(i)] = item;
    }

    @Override
    public Iterator<T> iterator() {
        // TODO Auto-generated method stub
        return new CircularIterator();
    }

    private class CircularIterator implements Iterator<T> {
        private int _current = -1;

        @Override
        public boolean hasNext() {
            // TODO Auto-generated method stub
            return _current < items.length - 1;
        }

        @Override
        public T next() {
            // TODO Auto-generated method stub
            _current++;
            return items[convert(_current)];
        }

    }

    /**
     * @param args
     */
    public static void main(String[] args) {
        // TODO Auto-generated method stub
        int mLength = 10;
        int mGetIndex = 5;
        CircularArray<Integer> mCircularArray = new CircularArray<Integer>(mLength);
        for (int i = 0; i < mLength; i++) {
            mCircularArray.set(i, i);
        }

        System.out.println("Get index:" + mGetIndex + " value:" + mCircularArray.get(mGetIndex));

        Iterator<Integer> mIter = mCircularArray.iterator();
        int mIndex = 0;
        while (mIter.hasNext()) {
            System.out.println("index:" + mIndex + " value: " + mIter.next());
            mIndex++;
        }

        int mShiftIndex = 4;
        mCircularArray.rotate(mShiftIndex);
        
        System.out.println("After rotate:" + mShiftIndex);

        mIter = mCircularArray.iterator();
        mIndex = 0;
        while (mIter.hasNext()) {
            System.out.println("index:" + mIndex + " value: " + mIter.next());
            mIndex++;
        }
        
        System.out.println("Get index:" + mGetIndex + " value:" + mCircularArray.get(mGetIndex));
    }
}
