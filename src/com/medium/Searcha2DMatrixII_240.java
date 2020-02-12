/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class Searcha2DMatrixII_240 {
    // 2020年2月13日
    // Use linear search
    // Time Complexity: O(m+n)
    // Space Complexity:O(1)
    // Runtime: 5 ms, faster than 99.88%
    // Memory Usage: 42.4 MB, less than 100.00%
    public boolean searchMatrix(int[][] matrix, int target) {
        if(matrix == null || matrix.length == 0) {
            return false;
        }
        
        int lengthRow = matrix.length, lengthCol = matrix[0].length;
        int indexRow = 0, indexCol = lengthCol - 1;
        
        while(indexRow < lengthRow && indexCol >= 0) {
            if(matrix[indexRow][indexCol] < target) {
                indexRow++;
            } else if(matrix[indexRow][indexCol] > target) {
                indexCol--;
            } else {
                return true;
            }
        }
        
        return false;
    }
}
