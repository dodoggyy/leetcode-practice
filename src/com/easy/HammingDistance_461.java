/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class HammingDistance_461 {

    // call bitCount API
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 32 MB, less than 100.00%
    public int hammingDistance(int x, int y) {
        return Integer.bitCount(x ^ y);
    }

    // without bitCount API
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 31.9 MB, less than 100.00%
    public int hammingDistance2(int x, int y) {
        int z = x ^ y;
        int mBitCount = 0;
        while (z != 0) {
            if ((z & 1) == 1) {
                mBitCount++;
            }
            z >>= 1;
        }
        return mBitCount;
    }
    
    // without bitCount API
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 31.9 MB, less than 100.00%
    public int hammingDistance3(int x, int y) {
        int z = x ^ y;
        int mBitCount = 0;
        while (z != 0) {
           z &= (z-1); // (z AND (z-1)) means remove LSB in 1 s'
           mBitCount++;
        }
        return mBitCount;
    }
}
