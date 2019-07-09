/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class AddBinary_67 {
    // Use sequential parsing 
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 2 ms, faster than 62.66% 
    // Memory Usage: 36.1 MB, less than 99.94%
    public String addBinary(String a, String b) {
        StringBuilder mBuilder = new StringBuilder();
        int mLengthA = a.length(), mLengthB = b.length();
        int mLength = Math.max(mLengthA, mLengthB);
        int mCarry = 0;
        for(int i = 0; i < mLength; i++) {
            int mTmpA = (mLengthA > i) ? a.charAt(mLengthA - i - 1) - '0' : 0;
            int mTmpB = (mLengthB > i) ? b.charAt(mLengthB - i - 1) - '0' : 0;
            mBuilder.insert(0, (mTmpA+mTmpB + mCarry)%2);
            mCarry = (mTmpA + mTmpB + mCarry) >> 1;
        }
        if(mCarry == 1) {
            mBuilder.insert(0, 1);
        }
        
        return mBuilder.toString();
//        return Integer.toBinaryString(Integer.parseInt(a, 2) + Integer.parseInt(b, 2));
    }
}
