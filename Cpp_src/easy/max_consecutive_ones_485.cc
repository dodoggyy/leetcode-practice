#include <iostream>
#include <vector>
using namespace std;

// Use loop judge:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 36 ms, faster than 84.61%
// Memory Usage: 11.8 MB, less than 72.73%
class Solution
{
public:
    int findMaxConsecutiveOnes(vector<int> &nums)
    {
        int result = 0;
        int tmp = 0;

        for (int num : nums)
        {
            if (num == 1)
            {
                tmp++;
            }
            else
            {
                result = max(result, tmp);
                tmp = 0;
            }
        }
        result = max(result, tmp);

        return result;
    }
};