#include <iostream>
#include <vector>
using namespace std;

// Use bit manipulation
// Time Complexity: O(n^2)
// Space Complexity:O(n)
// Runtime: 48 ms, faster than 76.35%
// Memory Usage: 13 MB, less than 80.00%
class Solution
{
public:
    int maxProduct(vector<string> &words)
    {
        int result = 0;
        vector<int> value(words.size(), 0);
        for (int i = 0; i < words.size(); i++)
        {
            for (char c : words[i])
            {
                value[i] |= (1 << (c - 'a'));
            }
        }
        for (int i = 0; i < words.size(); i++)
        {
            for (int j = i + 1; j < words.size(); j++)
            {
                if ((value[i] & value[j]) == 0)
                {
                    result = max(result, int(words[i].length() * words[j].length()));
                }
            }
        }
        return result;
    }
};