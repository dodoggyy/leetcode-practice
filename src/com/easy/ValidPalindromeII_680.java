/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ValidPalindromeII_680 {
    public static void main(String[] args) {
        // TODO Auto-generated method stub
        String s = "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga";

        System.out.println(validPalindrome(s));
    }
    
    // Runtime: 18 ms, faster than 89.92%
    // Memory Usage: 40.5 MB, less than 31.21%
    public static boolean validPalindrome(String s) {
        if (s.length() == 1) {
            return true;
        }
        if (s.length() > 50000) {
            return false;
        }
        int mLeftFlag = 0;
        int mRightFlag = s.length() - 1;
        
        while(mLeftFlag < mRightFlag) {
            if (s.charAt(mLeftFlag) != s.charAt(mRightFlag)) {
                return (isPalindrome(s, mLeftFlag+1, mRightFlag) | isPalindrome(s, mLeftFlag, mRightFlag-1));
            }
            mLeftFlag++;
            mRightFlag--;
        }
        
        return true;
    }
    
    public static boolean isPalindrome(String s, int mLeftFlag, int mRightFlag) {
        while (mLeftFlag < mRightFlag) {
            if (s.charAt(mLeftFlag++) != s.charAt(mRightFlag--)) {
                return false;
            }
        }
        return true;
    }
}
