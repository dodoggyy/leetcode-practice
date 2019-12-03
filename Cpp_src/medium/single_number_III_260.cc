#include <iostream>
#include <vector>
using namespace std;

// Use bit manipulation
// Time Complexity: O(2n) = O(n)
// Space Complexity:O(1)
// Runtime: 12 ms, faster than 95.69%
// Memory Usage: 9.7 MB, less than 100.00%
class Solution
{
public:
    vector<int> singleNumber(vector<int> &nums)
    {
        vector<int> result;
        int tmp_xor = 0;
        for (int num : nums)
        {
            tmp_xor ^= num;
        }
        int tmp_bit = (tmp_xor & (0 - tmp_xor));
        int result_1 = 0, result_2 = 0;
        for (int num : nums)
        {
            if (tmp_bit & num)
            {
                result_1 ^= num;
            }
            else
            {
                result_2 ^= num;
            }
        }
        result.push_back(result_1);
        result.push_back(result_2);

        return result;
    }
};