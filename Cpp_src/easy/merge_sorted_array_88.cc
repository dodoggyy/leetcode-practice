#include <iostream>
#include <vector>
using namespace std;

// Use conquer strategy without addition array:
// Time Complexity: O(m+n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 85.29%
// Memory Usage: 8.8 MB, less than 65.22%
class Solution
{
public:
    void merge(vector<int> &nums1, int m, vector<int> &nums2, int n)
    {
        int index_1 = m - 1;
        int index_2 = n - 1;
        int index_cur = m + n - 1;

        while (index_1 >= 0 && index_2 >= 0)
        {
            if (nums1[index_1] > nums2[index_2])
            {
                nums1[index_cur--] = nums1[index_1--];
            }
            else
            {
                nums1[index_cur--] = nums2[index_2--];
            }
        }
        while (index_2 >= 0)
        {
            nums1[index_cur--] = nums2[index_2--];
        }
    }
};