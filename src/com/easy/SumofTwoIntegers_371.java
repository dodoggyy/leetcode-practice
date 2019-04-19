/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class SumofTwoIntegers_371 {

    /**
     * @param args
     */
    public static void main(String[] args) {
        // TODO Auto-generated method stub

    }

    // recursive version
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 31.8 MB, less than 100.00%
    public int getSum(int a, int b) {
        // (a^b) | ((a & b) << 1);
        return (b == 0) ? a : getSum((a ^ b), ((a & b) << 1));
    }

    // iterative version
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 31.7 MB, less than 100.00%
    public int getSum2(int a, int b) {

        while (b != 0) {
            int mCarry = (a & b) << 1;
            a = (a ^ b);
            b = mCarry;
        }

        return a;
    }
}
