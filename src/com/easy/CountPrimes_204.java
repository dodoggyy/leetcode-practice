/**
 * 
 */
package com.easy;

import java.util.Arrays;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class CountPrimes_204 {
    // Use sieve of Eratosthenes
    // Time Complexity: O(nloglogn)
    // Space Complexity: O(n)
    // Runtime: 13 ms, faster than 48.29%
    // Memory Usage: 34.2 MB, less than 10.34%
    public int countPrimes(int n) {
        if (n <= 2) {
            return 0;
        }
        int mResult = 0;
        boolean[] bIsPrime = new boolean[n + 1];
        Arrays.fill(bIsPrime, true);
        bIsPrime[0] = false;
        bIsPrime[1] = false;
        for (int i = 2; i < n; i++) {
            if (!bIsPrime[i]) {
                continue;
            }
            ++mResult;
            for (int j = 2 * i; j < n; j += i) {
                bIsPrime[j] = false;
            }
        }
        return mResult;
    }

    // Use simple judgment under sqrt(n)
    // Time Complexity: O(n^1.5)
    // Space Complexity: O(1)
    // Runtime: 296 ms, faster than 15.77%
    // Memory Usage: 32.9 MB, less than 81.17%
    public int countPrimes2(int n) {
        if (n <= 2) {
            return 0;
        }
        int mResult = 1; // 2 is prime

        for (int i = 3; i < n; i++) {
            if (bIsPrime(i)) {
                mResult++;
            }
        }

        return mResult;
    }

    private boolean bIsPrime(int n) {
        if (n % 2 == 0) {
            return false;
        }
        for (int i = 3; i <= Math.sqrt(n); i += 2) {
            if (n % i == 0) {
                return false;
            }
        }
        return true;
    }
}
