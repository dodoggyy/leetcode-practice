/**
 * 
 */
package com.medium;

import java.util.Arrays;
import java.util.PriorityQueue;

/**
 * @author Chris Lin
 * @version 1.0
 */
public class KthSmallestElementinaSortedMatrix_378 {
    // 2020年2月12日
    // Use sort:
    // Time Complexity: O(nlogn)
    // Space Complexity:O(n)
    // Runtime: 7 ms, faster than 68.83%
    // Memory Usage: 48.1 MB, less than 8.11%
    public int kthSmallest(int[][] matrix, int k) {
        int length = matrix.length;
        int[] tmp = new int[length * length];
        for (int i = 0; i < length * length; i++) {
            tmp[i] = matrix[i / length][i % length];
        }
        Arrays.sort(tmp);

        return tmp[k - 1];
    }

    // 2020年2月12日
    // Use priority queue:
    // Time Complexity: O(nlogn*logk)
    // Space Complexity:O(k)
    // Runtime: 22 ms, faster than 22.09%
    // Memory Usage: 46.5 MB, less than 13.51%
    public int kthSmallest2(int[][] matrix, int k) {
        int length = matrix.length;
        PriorityQueue<Integer> mQueue = new PriorityQueue<>();
        int target = length * length - k + 1;
        for (int i = 0; i < length * length; i++) {
            mQueue.add(matrix[i / length][i % length]);
            if (mQueue.size() > target) {
                mQueue.poll();
            }
        }
        return mQueue.peek();
    }

    // 2020年2月12日
    // Use binary search:
    // Time Complexity: O(nlogx)
    // Space Complexity:O(1)
    // x : difference between max and min
    // Runtime: 1 ms, faster than 87.43%
    // Memory Usage: 51.8 MB, less than 5.41%
    public int kthSmallest3(int[][] matrix, int k) {
        int length = matrix.length;
        int low = matrix[0][0], high = matrix[length - 1][length - 1];

        while (low <= high) {
            int mid = low + (high - low) / 2;
            int cnt = countIndex(matrix, mid);
            if (cnt < k) {
                low = mid + 1;
            } else {
                high = mid - 1;
            }
        }
        return low;
    }

    public int countIndex(int[][] matrix, int target) {
        int length = matrix.length, i = length - 1, j = 0, cnt = 0;
        while (i >= 0 && j < length) {
            if (matrix[i][j] <= target) {
                cnt += i + 1;
                j++;
            } else {
                i--;
            }
        }
        return cnt;
    }
}
