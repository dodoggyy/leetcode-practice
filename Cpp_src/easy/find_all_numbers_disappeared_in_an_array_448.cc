#include <iostream>
#include <vector>
using namespace std;

// Use index mapping:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 112 ms, faster than 92.30%
// Memory Usage: 14.9 MB, less than 80.00%
class Solution
{
public:
    vector<int> findDisappearedNumbers(vector<int> &nums)
    {
        vector<int> result;
        for (int i = 0; i < nums.size(); i++)
        {
            int index = abs(nums[i]) - 1;
            if (nums[index] > 0)
            {
                nums[index] *= -1;
            }
        }

        for (int i = 0; i < nums.size(); i++)
        {
            if (nums[i] > 0)
            {
                result.push_back(i + 1);
            }
        }
        return result;
    }
};