/**
 * 
 */
package com.easy;

import java.util.Arrays;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ReverseVowelsOfAString_345 {
    public static void main(String[] args) {
        // TODO Auto-generated method stub
        String s = "hello";
        char[] mResult = s.toCharArray();

        System.out.println(reverseVowels(s));
    }

    // Runtime: 6 ms, faster than 61.71%
    // Memory Usage: 37.8 MB, less than 66.79%
    public static String reverseVowels(String s) {
        if (s.length() <= 1) {
            return s;
        }

        // Cannot use char because it cannot use Arrays.asList() for primitive arrays
        Character[] mVowels = { 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U' };
        char[] mResult = new char[s.length()];
        char mTmpLeft, mTmpRight;
        int mLeftFlag = 0;
        int mRightFlag = mResult.length - 1;

        while (mLeftFlag <= mRightFlag) {
            mTmpLeft = s.charAt(mLeftFlag);
            mTmpRight = s.charAt(mRightFlag);
            if(!Arrays.asList(mVowels).contains(mTmpLeft)) {
                mResult[mLeftFlag] = mTmpLeft;
//                System.out.printf("mLeftFlag:%d mTmpLeft:%c\n",mLeftFlag,mTmpLeft);
                mLeftFlag++;
            } else if(!Arrays.asList(mVowels).contains(mTmpRight)) {
                mResult[mRightFlag] = mTmpRight;
//                System.out.printf("mRightFlag:%d mTmpRight:%c\n",mRightFlag,mTmpRight);
                mRightFlag--;
            } else {
                mResult[mLeftFlag] = mTmpRight;
                mResult[mRightFlag] = mTmpLeft;
                mLeftFlag++;
                mRightFlag--;
//                System.out.printf("mLeftFlag:%d mTmpLeft:%c  mRightFlag:%d mTmpRight:%c\n",mLeftFlag,mTmpLeft,mRightFlag,mTmpRight);
            }
        }

        return String.valueOf(mResult);
    }
    
    // Runtime: 2 ms, faster than 99.62%
    // Memory Usage: 38.3 MB, less than 61.37%
    public static String reverseVowels2(String s) {
        if (s.length() <= 1) {
            return s;
        }
        int mLeftFlag = 0; 
        int mRightFlag = s.length() - 1; 
        char[] mResult = s.toCharArray();
        
        while (mLeftFlag < mRightFlag) {
           if(!bIsVowel(mResult[mLeftFlag])) {
               mLeftFlag++;
               continue;
           }
           if (!bIsVowel(mResult[mRightFlag])) {
               mRightFlag--;
               continue;
           }

           //swap vowel if find
           char tmp = mResult[mLeftFlag];
           mResult[mLeftFlag] = mResult[mRightFlag];
           mResult[mRightFlag] = tmp;
           
           mLeftFlag++;
           mRightFlag--;
        }
        
        return String.valueOf(mResult);
    }
    
    public static boolean bIsVowel(char c) {
        if(c=='a' || c=='e' || c=='i' || c=='o' || c=='u' ||
                c=='A' || c=='E' || c=='I' || c=='O' || c=='U')
            return true;
        return false;
    }
}
