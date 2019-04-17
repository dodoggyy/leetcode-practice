/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class PowerofTwo_231 {
    public static void main(String[] args) {
        int n = -2147483648;
        isPowerOfTwo(n);
    }
    
    // Runtime: 1 ms, faster than 91.87%
    // Memory Usage: 32.2 MB, less than 100.00%
    public static boolean isPowerOfTwo(int n) {
        //System.out.println(Integer.bitCount(n));
        return ((Integer.bitCount(n) != 1) || (n < 0))?false:true;
    }
}
