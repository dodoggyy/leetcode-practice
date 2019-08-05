/**
 * 
 */
package com.medium;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class WordSearch_79 {
    // 2019年8月5日
    // Use DFS
    // Time Complexity: O(m*n*4)
    // Space Complexity:O(m*n + l)
    // m: board row length, n: board column length, l: word length
    // Runtime: 4 ms, faster than 89.53%
    // Memory Usage: 38.4 MB, less than 98.42%
    public boolean exist(char[][] board, String word) {
        if (board == null || word == null || board.length == 0) {
            return false;
        }
        int mLengthRow = board.length, mLengthCol = board[0].length;
        for (int i = 0; i < mLengthRow * mLengthCol; i++) {
            if (search(board, word, i / mLengthCol, i % mLengthCol, 0)) {
                return true;
            }
        }
        return false;
    }

    private boolean search(char[][] board, String word, int mRow, int mCol, int mLengthPath) {
        if (mRow < 0 || mCol < 0 || mRow == board.length || mCol == board[0].length
                || word.charAt(mLengthPath) != board[mRow][mCol]) {
            return false;
        }
        if (mLengthPath == word.length() - 1) {
            return true;
        }

        char mCurrentChar = board[mRow][mCol];
        board[mRow][mCol] = '#';
        boolean bIsFound = search(board, word, mRow + 1, mCol, mLengthPath + 1)
                || search(board, word, mRow - 1, mCol, mLengthPath + 1)
                || search(board, word, mRow, mCol + 1, mLengthPath + 1)
                || search(board, word, mRow, mCol - 1, mLengthPath + 1);
        board[mRow][mCol] = mCurrentChar;
        return bIsFound;
    }
}
