#include <iostream>
#include <vector>
using namespace std;

// Use index copy:
// Time Complexity: O(m*n)
// Space Complexity:O(m*n)
// Runtime: 32 ms, faster than 92.71%
// Memory Usage: 11.6 MB, less than 80.00%
class Solution
{
public:
    vector<vector<int>> matrixReshape(vector<vector<int>> &nums, int r, int c)
    {
        int origin_r = nums.size(), origin_c = nums[0].size();
        vector<vector<int>> result(r, vector<int>(c, 0));

        if (origin_r == 0 || origin_r * origin_c != r * c)
        {
            return nums;
        }

        int total_element = origin_r * origin_c;
        for (int i = 0; i < total_element; i++)
        {
            result[i / c][i % c] = nums[i / origin_c][i % origin_c];
        }

        return result;
    }
};