/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class LicenseKeyFormatting_482 {
    // 2019年7月22日
    // Use replacement & comparison
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 138 ms, faster than 14.36%
    // Memory Usage: 38.6 MB, less than 80.73%
    public String licenseKeyFormatting(String S, int K) {
        StringBuilder mResult = new StringBuilder();

        S = S.replace("-", "").toUpperCase();
        int mTmpDash = S.length() % K;
        for (int i = S.length() - 1; i >= 0; i--) {
            mResult.insert(0, S.charAt(i));
            if ((i - mTmpDash) % K == 0 && i != 0) {
                mResult.insert(0, "-");
            }
        }

        return mResult.toString();
    }

    // 2019年7月22日
    // Optimize replacement & comparison
    // Time Complexity: O(n)
    // Space Complexity:O(n)
    // Runtime: 20 ms, faster than 47.55%
    // Memory Usage: 37.3 MB, less than 99.89%
    public String licenseKeyFormatting2(String S, int K) {
        StringBuilder mResult = new StringBuilder();

        for (int i = S.length() - 1; i >= 0; i--) {
            if (S.charAt(i) != '-') {
                if (mResult.length() % (K + 1) == K) {
                    mResult.append('-');
                }
                mResult.append(S.charAt(i));
            }
        }

        return mResult.reverse().toString().toUpperCase();
    }
}
