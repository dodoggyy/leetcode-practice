#include <iostream>
using namespace std;

// Use array count:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 99.40%
// Memory Usage: 9.1 MB, less than 86.00%
class Solution
{
public:
    bool isIsomorphic(string s, string t)
    {
        int index_s[256] = {0}, index_t[256] = {0};
        for (int i = 0; i < s.size(); i++)
        {
            if (index_s[s[i]] != index_t[t[i]])
            {
                return false;
            }
            index_s[s[i]] = i + 1;
            index_t[t[i]] = i + 1;
        }

        return true;
    }
};