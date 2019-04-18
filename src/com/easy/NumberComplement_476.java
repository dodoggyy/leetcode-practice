/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class NumberComplement_476 {

    /**
     * @param args
     */
    public static void main(String[] args) {
        // TODO Auto-generated method stub
        int num = 0;
        System.out.println(findComplement(num));
    }

    // with HighestOneBit API
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 32 MB, less than 100.00%
    public static int findComplement(int num) {
        System.out.println(Integer.highestOneBit(num));
        if (num == 0) {
            return 1;
        }
        int mHighBit = Integer.highestOneBit(num);
        return (((mHighBit << 1) - 1) ^ num); // e.g. 0xFFFF ^ 0x1111 = 0x1111 complement
    }

    // without HighestOneBit API
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 31.9 MB, less than 100.00%
    public static int findComplement2(int num) {
        if (num == 0) {
            return 1;
        }
        int mMask = (1 << 30);
        while ((num & mMask) == 0) {
            mMask >>= 1;
        }
        mMask = (mMask << 1) - 1;
        return (num ^ mMask);
    }
}
