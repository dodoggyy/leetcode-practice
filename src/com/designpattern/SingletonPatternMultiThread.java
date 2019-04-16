/**
 * 
 */
package com.designpattern;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class SingletonPatternMultiThread {

    private static SingletonPatternMultiThread instance = null;

    private SingletonPatternMultiThread() {

    }

    public static SingletonPatternMultiThread getInstance() {
        if (instance == null) { // Double-check Locking
            synchronized (SingletonPatternMultiThread.class) { // sync
                if (instance == null) {
                    instance = new SingletonPatternMultiThread();
                }
            }
        }
        return instance;
    }

    /**
     * @param args
     */
    public static void main(String[] args) {
        // TODO Auto-generated method stub

    }

}
