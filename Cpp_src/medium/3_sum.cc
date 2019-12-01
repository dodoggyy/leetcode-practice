#include <algorithm>
#include <iostream>
#include <vector>
using namespace std;

// Reduce 2-sum problem and use 2 pointers
// Time Complexity: O(n^2)
// Space Complexity:O(1)
// Runtime: 92 ms, faster than 94.17%
// Memory Usage: 15.7 MB, less than 50.00%
class Solution
{
public:
    vector<vector<int>> threeSum(vector<int> &nums)
    {
        vector<vector<int>> result;
        if (nums.empty() || nums.size() < 3)
        {
            return result;
        }
        sort(nums.begin(), nums.end());

        for (int i = 0; i < nums.size() - 2; i++)
        {
            if (i > 0 && nums[i] == nums[i - 1])
            {
                continue;
            }
            int complement = 0 - nums.at(i);
            int index_front = i + 1;
            int index_tail = nums.size() - 1;
            while (index_front < index_tail)
            {
                int sum = nums[index_front] + nums[index_tail];
                if (sum == complement)
                {
                    vector<int> tmp = {nums[i], nums[index_front], nums[index_tail]};
                    result.push_back(tmp);
                    while (index_front < index_tail && nums[index_front] == nums[index_front + 1])
                    {
                        index_front++;
                    }
                    while (index_front < index_tail && nums[index_tail] == nums[index_tail - 1])
                    {
                        index_tail--;
                    }
                    index_front++;
                    index_tail--;
                }
                else if (sum < complement)
                {
                    index_front++;
                }
                else
                {
                    index_tail--;
                }
            }
        }

        return result;
    }
};