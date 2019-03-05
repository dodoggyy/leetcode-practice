/**
 * 
 */
package com.easy;

/**
 * @author Chris Lin
 * @version 1.0
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
            tmp[mIndexTmp++] = (nums1[mIndexM] > nums2[mIndexN])?nums2[mIndexN++]:nums1[mIndexM++];
        }
        while(mIndexM < m) { // add remain data from nums1 to tmp
            tmp[mIndexTmp++] = nums1[mIndexM++];
        }
        while(mIndexN < n) { // add remain data from nums2 to tmp
            tmp[mIndexTmp++] = nums2[mIndexN++];
        }

        for (int i = 0; i < m+n; i++) { // move entire merged list to nums1
            nums1[i] = tmp[i];
        }
    }

}
