#include <iostream>
#include <vector>
using namespace std;

// Use bit manipulation
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 16 ms, faster than 67.73%
// Memory Usage: 9.5 MB, less than 100.00%
class Solution
{
public:
    int singleNumber(vector<int> &nums)
    {
        int result = 0;
        for (int num : nums)
        {
            result ^= num;
        }
        return result;
    }
};