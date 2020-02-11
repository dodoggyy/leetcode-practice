#include <iostream>
using namespace std;

// Use Expand Around Center
// Time Complexity: O(n^2)
// Space Complexity:O(1)
// Runtime: 12 ms, faster than 58.78%
// Memory Usage: 15.5 MB, less than 16.00%
class Solution
{
public:
    int result = 0;
    int countSubstrings(string s)
    {
        for (int i = 0; i < s.size(); i++)
        {
            helper(s, i, i);
            helper(s, i, i + 1);
        }

        return result;
    }

    void helper(string s, int index_left, int index_right)
    {
        while (index_left >= 0 && index_right < s.size() && s[index_left] == s[index_right])
        {
            index_left--;
            index_right++;
            result++;
        }
    }
};