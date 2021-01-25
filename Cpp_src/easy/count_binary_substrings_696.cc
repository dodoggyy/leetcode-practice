#include <iostream>
using namespace std;

// Use Group By Character
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 44 ms, faster than 33.53%
// Memory Usage: 13 MB, less than 25.00%
class Solution
{
public:
    int countBinarySubstrings(string s)
    {
        int result = 0, index = 0, count = 1;
        int group[s.size()];

        for (int i = 1; i < s.size(); i++)
        {
            if (s[i] != s[i - 1])
            {
                group[index++] = count;
                count = 1;
            }
            else
            {
                count++;
            }
        }
        group[index++] = count;

        for (int i = 1; i < index; i++)
        {
            result += min(group[i], group[i - 1]);
        }
        return result;
    }
};

// Use Linear Scan
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 36 ms, faster than 91.12%
// Memory Usage: 12.8 MB, less than 25.00%
class Solution2
{
public:
    int countBinarySubstrings(string s)
    {
        int result = 0;
        int pre_cnt = 0, cur_cnt = 1;

        for (int i = 1; i < s.size(); i++)
        {
            if (s[i] != s[i - 1])
            {
                result += min(pre_cnt, cur_cnt);
                pre_cnt = cur_cnt;
                cur_cnt = 1;
            }
            else
            {
                cur_cnt++;
            }
        }
        result += min(pre_cnt, cur_cnt);

        return result;
    }
};