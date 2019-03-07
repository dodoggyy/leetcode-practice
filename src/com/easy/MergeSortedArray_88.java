/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.1
 */
public class MergeSortedArray_88 {

    // Runtime: 2 ms, faster than 100.00% of Java
    // Memory Usage: 37.3 MB, less than 72.80% of Java
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        int mIndexStart = 0;
        int mIndexM = 0;
        int mIndexN = 0;
        int mIndexTmp = 0;
        int[] tmp = new int[m + n];

        while (mIndexM < m && mIndexN < n) { // ref list merge method
            tmp[mIndexTmp++] = (nums1[mIndexM] > nums2[mIndexN]) ? nums2[mIndexN++] : nums1[mIndexM++];
        }
        while (mIndexM < m) { // add remain data from nums1 to tmp
            tmp[mIndexTmp++] = nums1[mIndexM++];
        }
        while (mIndexN < n) { // add remain data from nums2 to tmp
            tmp[mIndexTmp++] = nums2[mIndexN++];
        }

        for (int i = 0; i < m + n; i++) { // move entire merged list to nums1
            nums1[i] = tmp[i];
        }
    }

    // Runtime: 2 ms, faster than 100.00% of Java
    // Memory Usage: 37.5 MB, less than 6.85% of Java
    public void merge2(int[] nums1, int m, int[] nums2, int n) {
        int mIndexM = (m - 1);
        int mIndexN = (n - 1);
        int mIndexTmp = (m + n - 1);

        while (mIndexM >= 0 || mIndexN >= 0) {
            if (mIndexM < 0) {
                nums1[mIndexTmp--] = nums2[mIndexN--];
            } else if (mIndexN < 0) {
                nums1[mIndexTmp--] = nums1[mIndexM--];
            } else if (nums1[mIndexM] >= nums2[mIndexN]) {
                nums1[mIndexTmp--] = nums1[mIndexM--];
            } else {
                nums1[mIndexTmp--] = nums2[mIndexN--];
            }
        }
    }
}
