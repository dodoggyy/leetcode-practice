/**
 * 
 */
package com.easy;

import java.util.HashMap;
import java.util.Map;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ReverseBits_190 {

    private Map<Byte, Integer> mMap = new HashMap<Byte, Integer>();

    // you need treat n as an unsigned value
    // 1 times version
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 28.1 MB, less than 55.35%
    public int reverseBits(int n) {
        int mResult = 0;

        for (int i = 0; i < 32; i++) {
            mResult <<= 1;
            mResult |= (n & 1);
            n >>= 1;
        }

        return mResult;
    }

    // separate by byte version
    // Runtime: 3 ms, faster than 8.15%
    // Memory Usage: 28.1 MB, less than 54.98%
    public int reverseBits2(int n) {
        int mResult = 0;
        for (int i = 0; i < 4; i++) {
            mResult <<= 8;
            mResult |= reverseByte((byte) (n & 0xFF));
            n >>= 8;
        }

        return mResult;
    }

    private int reverseByte(byte mByte) {
        if (mMap.containsKey(mByte)) {
            return mMap.get(mByte);
        }
        int mResult = 0;
        byte mTmpByte = mByte;

        for (int i = 0; i < 8; i++) {
            mResult <<= 1;
            mResult |= (mTmpByte & 1);
            mTmpByte >>= 1;
        }
        mMap.put(mByte, mResult);

        return mResult;
    }
}
