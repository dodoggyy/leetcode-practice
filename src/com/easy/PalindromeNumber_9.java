/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 * {@link https://github.com/dodoggyy/AlgorithmPractice/blob/master/src/com/practice/Palindrome.java}.
 */
public class PalindromeNumber_9 {

    // Runtime: 6 ms, faster than 100.00%
    // Memory Usage: 35.2 MB, less than 100.00%
    public boolean isPalindrome(int x) {
        if(x < 0) {
            return false;
        }
        
        int mReverse = 0, mCopy = x;
        while(x != 0) {
            mReverse = (mReverse*10) + (x%10);
            x /= 10;
        }
        return (mReverse == mCopy);
    }
}
