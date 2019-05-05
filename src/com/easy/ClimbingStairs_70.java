/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class ClimbingStairs_70 {
    
    // Iterative version
    // Space complexity: O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 31.8 MB, less than 100.00%
    public int climbStairs(int n) {
        if(n <= 2) {
            return n;
        }
        int fibNminus2 = 1, fibNminus1 = 2;
            int fibN = fibNminus2 + fibNminus1;
            fibNminus2 = fibNminus1;
            fibNminus1 = fibN;
        
        return fibNminus1;
    }
    
    // Recursive version
    // Space complexity: O(n)
    // Time Limit Exceeded
    public int climbStairs2(int n) {
        if(n <= 2) {
            return n;
        }
        return climbStairs2(n-1) + climbStairs2(n-2);
    }
    
}
