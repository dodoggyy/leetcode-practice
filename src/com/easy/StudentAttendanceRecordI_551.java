/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class StudentAttendanceRecordI_551 {
    // 2019年7月28日
    // Use iterative comparison
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00%
    // Memory Usage: 34.5 MB, less than 100.00%
    public boolean checkRecord(String s) {
        int mCountLate = 0, mCountAbsent = 0;
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == 'A') {
                if (++mCountAbsent > 1) {
                    return false;
                }

            } else if (s.charAt(i) == 'L') {
                if (++mCountLate > 2 && s.charAt(i - 1) == 'L' && s.charAt(i - 2) == 'L') {
                    return false;
                }
            }
        }
        return true;
    }

    // 2019年7月28日
    // Use regular expression
    // Time Complexity: O(n)
    // Space Complexity:O(1)
    // Runtime: 4 ms, faster than 8.28%
    // Memory Usage: 35.1 MB, less than 100.00%
    public boolean checkRecord2(String s) {
        return !s.matches(".*A.*A.*|.*LLL.*");
    }
}
