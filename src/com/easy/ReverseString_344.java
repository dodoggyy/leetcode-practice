/**
 * 
 */
package com.easy;

import java.util.Stack;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ReverseString_344 {

    // Swap head and tail
    // Time Complexity: O(n/2) = O(n)
    // Space Complexity:O(1)
    // Runtime: 1 ms, faster than 100.00%
    // Memory Usage: 43.9 MB, less than 86.53%
    public void reverseString(char[] s) {
        for (int i = 0; i < (s.length / 2); i++) {
            swap(s, i, s.length - i - 1);
        }
    }

    private void swap(char[] s, int mIndex1, int mIndex2) {
        char tmp = s[mIndex1];
        s[mIndex1] = s[mIndex2];
        s[mIndex2] = tmp;
    }
    
    // XOR swap
    // a = a^b
    // b = a^b
    // a = a^b
    private void swap2(char[] s, int mIndex1, int mIndex2) {
        s[mIndex1] = (char) (s[mIndex1]^ s[mIndex2]);
        s[mIndex2] = (char) (s[mIndex1]^ s[mIndex2]);
        s[mIndex1] = (char) (s[mIndex1]^ s[mIndex2]);
    }
    
    // Use stack
    // Time Complexity: O(2n) = O(n)
    // Space Complexity:O(n)
    // Runtime: 9 ms, faster than 60.04% 
    // Memory Usage: 45.4 MB, less than 84.46%
    public void reverseString2(char[] s) {
      Stack<Character> mStack = new Stack<>();
      for(int i = 0; i < s.length; i++) {
          mStack.push(s[i]);
      }
      for(int i = 0; i < s.length; i++) {
          s[i] = mStack.pop();
      }
    }
}
