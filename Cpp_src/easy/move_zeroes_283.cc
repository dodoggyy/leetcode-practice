#include <iostream>
#include <vector>
using namespace std;

// Use Index to record non-zero
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 16 ms, faster than 65.42%
// Memory Usage: 9.6 MB, less than 61.11%
class Solution
{
public:
    void moveZeroes(vector<int> &nums)
    {
        int index_non_zero = 0;

        for (int i = 0; i < nums.size(); i++)
        {
            if (nums[i] != 0)
            {
                swap(nums, i, index_non_zero++);
            }
        }
    }

    void swap(vector<int> &nums, int index1, int index2)
    {
        int tmp = nums[index1];
        nums[index1] = nums[index2];
        nums[index2] = tmp;
    }
};