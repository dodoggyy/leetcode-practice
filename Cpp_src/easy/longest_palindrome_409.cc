#include <iostream>
using namespace std;

// Use greedy algorithm:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 8.9 MB, less than 73.33%
class Solution
{
public:
    int longestPalindrome(string s)
    {
        int result = 0;
        int count[128] = {0};

        for (char c : s)
        {
            count[c]++;
        }

        for (int cnt : count)
        {
            result += (cnt >> 1) << 1;
            if ((result & 1) == 0 && (cnt & 1))
            {
                result++;
            }
        }

        return result;
    }
};