/**
 * 
 */
package com.codility;

import java.util.Stack;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class Brackets_L7 {
    // 2020年3月9日
    // Use stack:
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    public int solution(String S) {
        // write your code in Java SE 8
        if (S == null) {
            return 0;
        }
        Stack<Character> stack = new Stack<>();
        for (char c : S.toCharArray()) {
            switch (c) {
            case '{':
            case '(':
            case '[':
                stack.push(c);
                break;
            case '}':
                if (stack.isEmpty() || stack.pop() != '{') {
                    return 0;
                }
                break;
            case ')':
                if (stack.isEmpty() || stack.pop() != '(') {
                    return 0;
                }
                break;
            case ']':
                if (stack.isEmpty() || stack.pop() != '[') {
                    return 0;
                }
                break;
            default:
                return 0;
            }
        }
        if (!stack.isEmpty()) {
            return 0;
        }
        return 1;
    }
}
