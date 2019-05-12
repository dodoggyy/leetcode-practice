/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class RotateImage_48 {
    
    // Direct rotate
    // Time Complexity: O(n^2)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00% 
    // Memory Usage: 35.3 MB, less than 93.70%
    public void rotate(int[][] matrix) {
        int n = matrix.length;
        for(int i = 0; i < n/2; i++) {
            for(int j = i; j < (n - 1 - i) ; j++) {
                int tmp = matrix[i][j];
                matrix[i][j] = matrix[n-1-j][i];
                matrix[n-1-j][i] = matrix[n-1-i][n-1-j];
                matrix[n-1-i][n-1-j] = matrix[j][n-1-i];
                matrix[j][n-1-i] = tmp;
            }
        }
    }
    
    // Diagonal for axis flip and X-axis mid-line flip
    // Time Complexity: O(2*n^2) = O(n^2)
    // Space Complexity:O(1)
    // Runtime: 0 ms, faster than 100.00% 
    // Memory Usage: 35.2 MB, less than 93.70%
    public void rotate2(int[][] matrix) {
        int n = matrix.length;
        for(int i = 0; i < n-1; i++) {
            for(int j = 0; j < n-i; j++) {
                int tmp = matrix[i][j];
                matrix[i][j] = matrix[n-1-j][n-1-i];
                matrix[n-1-j][n-1-i] = tmp;
            }
        }
        for(int i = 0; i < n/2; i++) {
            for(int j = 0; j < n;j++) {
                int tmp = matrix[i][j];
                matrix[i][j] = matrix[n-1-i][j];
                matrix[n-1-i][j] = tmp;
            }
        }
    }
    
}
