/**
 * 
 */
package com.codility;

import java.util.Stack;
/**
 * @author Chris Lin
 * @version 1.0
 */
public class Nesting_L7 {
    // 2020年3月9日
    // Use stack:
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    public int solution(String S) {
        // write your code in Java SE 8
        if(S == null) {
            return 0;
        }
        Stack<Character> stack = new Stack<>();
        for(char c: S.toCharArray()) {
            if(c == '(') {
                stack.push('(');
            } else if(c == ')') {
                if(stack.isEmpty()) {
                    return 0;
                }
                stack.pop();
            }
        }
        if(!stack.isEmpty()) {
            return 0;
        }
        return 1;
    }
}
