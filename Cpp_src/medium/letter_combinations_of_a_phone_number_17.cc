#include <algorithm>
#include <iostream>
#include <vector>
using namespace std;

// Use DFS
// Time Complexity: O(4^n)
// Space Complexity:O(4^n + n)
// Runtime: 4 ms, faster than 60.35%
// Memory Usage: 8.5 MB, less than 88.57%
class Solution
{
public:
    vector<string> letterCombinations(string digits)
    {
        vector<string> result;
        if (digits.size() == 0)
        {
            return result;
        }

        dfs(digits, 0, "", result);
        return result;
    }

private:
    string KEYS[10] = {"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"};

    void dfs(string digits, int offset, string tmp, vector<string> &result)
    {
        if (offset == digits.size())
        {
            result.push_back(tmp);
            return;
        }

        string letters = KEYS[digits[offset] - '0'];
        for (int i = 0; i < letters.size(); i++)
        {
            dfs(digits, offset + 1, tmp + letters[i], result);
        }
    }
};

// Use BFS
// Time Complexity: O(4^n)
// Space Complexity:O(2*(4^n))
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 8.4 MB, less than 100.00%
class Solution2
{
public:
    vector<string> letterCombinations(string digits)
    {
        vector<string> result;
        if (digits.empty())
        {
            return result;
        }
        result.push_back("");
        for (char digit : digits)
        {
            vector<string> tmp;
            for (string s : result)
            {
                for (char c : KEYS[digit - '0'])
                {
                    tmp.push_back(s + c);
                }
            }
            result.swap(tmp);
        }
        return result;
    }

private:
    string KEYS[10] = {"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"};
};