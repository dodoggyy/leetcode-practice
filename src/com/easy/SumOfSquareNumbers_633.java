/**
 * 
 */
package com.easy;

import java.util.HashMap;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class SumOfSquareNumbers_633 {
    
    // Runtime: 74 ms, faster than 11.32%
    // Memory Usage: 39.3 MB, less than 0.93%
    public boolean judgeSquareSum(int c) {
        int tmp = 0;
        tmp = (int)Math.sqrt(c);
        HashMap<Integer, Integer> mMap = new HashMap<>();
        
        for(int i = 0; i <= tmp; i++) {
            mMap.put((c - i*i),  i);
        }
        for(int i = 0; i <= tmp; i++) {
            if(mMap.containsKey(i*i)) {
                return true;
            }
        }
        
        
        return false;
    }

    // Runtime: 7 ms, faster than 92.77% 
    // Memory Usage: 26 MB, less than 51.85%
    public boolean judgeSquareSum2(int c) {
        int a = 0;
        int b = (int)Math.sqrt(c);
        int tmp = 0;
        
        if(c < 0) {
            return false;
        }
        
        while(a <= b) {
            tmp = a*a + b*b;
            
            if(tmp < c ) {
                a++;
            } else if(tmp > c) {
                b--;
            }else if(tmp == c ) {
                return true;
            }
        }
        return false;
    }
}
