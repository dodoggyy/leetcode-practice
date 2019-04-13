/**
 * 
 */
package com.easy;

import java.util.Arrays;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class AssignCookies_455 {

    /**
     * @param args
     */
    public static void main(String[] args) {
        // TODO Auto-generated method stub
        int[] g = { 10, 9, 8, 7 };
        int[] s = { 5, 6, 7, 8 };

        System.out.println(findContentChildren(g, s));
    }

    // Runtime: 8 ms, faster than 99.87%
    // Memory Usage: 40.8 MB, less than 12.93%
    public static int findContentChildren(int[] g, int[] s) {
        // g means greed factor
        // s means cookie size
        if ((g.length == 0) || (s.length == 0)) {
            return 0;
        }

        Arrays.sort(g);
        Arrays.sort(s);

        int mIndexG = 0, mIndexS = 0; //, mResult = 0;

        while ((mIndexG < g.length) && (mIndexS < s.length)) {
            if (g[mIndexG] <= s[mIndexS]) {
                //mResult++;
                mIndexG++;
            }
            mIndexS++;
        }

        return mIndexG;//mResult;
    }
}
