#include <iostream>
#include <vector>
using namespace std;

// Use Boyer-Moore vote:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 16 ms, faster than 96.51%
// Memory Usage: 8.8 MB, less than 100.00%
class Solution
{
public:
    int majorityElement(vector<int> &nums)
    {
        int result = nums[0], cnt = 0;
        for (int num : nums)
        {
            if (num == result)
            {
                cnt++;
                if (cnt > nums.size())
                {
                    return result;
                }
            }
            else
            {
                cnt--;
                if (cnt == 0)
                {
                    result = num;
                    cnt = 1;
                }
            }
        }

        return result;
    }
};