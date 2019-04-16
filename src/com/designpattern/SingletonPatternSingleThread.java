/**
 * 
 */
package com.designpattern;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class SingletonPatternSingleThread {
    private static SingletonPatternSingleThread instance = null;

    private SingletonPatternSingleThread() {

    }

    public SingletonPatternSingleThread getInstance() {
        if (instance == null) {
            instance = new SingletonPatternSingleThread();
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
