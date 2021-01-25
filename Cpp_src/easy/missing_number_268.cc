#include <algorithm>
#include <iostream>
#include <vector>
using namespace std;

// Use bit manipulation
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 32 ms, faster than 21.72%
// Memory Usage: 9.7 MB, less than 98.04%
class Solution
{
public:
    int missingNumber(vector<int> &nums)
    {
        int result = nums.size();
        for (int i = 0; i < nums.size(); i++)
        {
            result ^= i;
            result ^= nums[i];
        }
        return result;
    }
};

// Use sum
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 24 ms, faster than 82.06%
// Memory Usage: 9.9 MB, less than 66.67%
class Solution2
{
public:
    int missingNumber(vector<int> &nums)
    {
        int sum = nums.size() * (nums.size() + 1) / 2;
        for (int num : nums)
        {
            sum -= num;
        }
        return sum;
    }
};

// Use binary search
// Time Complexity: O(nlogn)
// Space Complexity:O(1)
// Runtime: 32 ms, faster than 21.72%
// Memory Usage: 9.9 MB, less than 72.55%
class Solution3
{
public:
    int missingNumber(vector<int> &nums)
    {
        sort(nums.begin(), nums.end());
        int index_left = 0, index_right = nums.size();
        while (index_left < index_right)
        {
            int index_mid = index_left + (index_right - index_left) / 2;
            if (nums[index_mid] > index_mid)
            {
                index_right = index_mid;
            }
            else
            {
                index_left = index_mid + 1;
            }
        }
        return index_left;
    }
};