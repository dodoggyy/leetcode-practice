#include <iostream>
#include <vector>
using namespace std;

// Use horizontal scanning
// Time Complexity: O(S)
// Space Complexity:O(1)
// S is the sum of all characters in all strings
// Runtime: 4 ms, faster than 95.23%
// Memory Usage: 9 MB, less than 20.97%
class Solution
{
public:
    string longestCommonPrefix(vector<string> &strs)
    {
        if (strs.empty())
        {
            return "";
        }
        string result = strs[0];
        for (int i = 1; i < strs.size(); i++)
        {
            int index = 0;
            while (index < min(result.size(), strs[i].size()) && result[index] == strs[i][index])
            {
                index++;
            }
            result = result.substr(0, index);
        }
        return result;
    }
};

// Use vertical scanning
// Time Complexity: O(S)
// Space Complexity:O(1)
// S is the sum of all characters in all strings
// Runtime: 4 ms, faster than 95.23%
// Memory Usage: 8.9 MB, less than 77.42%
class Solution2
{
public:
    string longestCommonPrefix(vector<string> &strs)
    {
        if (strs.empty())
        {
            return "";
        }
        for (int i = 0; i < strs[0].length(); i++)
        {
            char c = strs[0][i];
            for (int j = 1; j < strs.size(); j++)
            {
                if (i == strs[j].length() || c != strs[j][i])
                {
                    return strs[0].substr(0, i);
                }
            }
        }
        return strs[0];
    }
};