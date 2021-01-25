#include <iostream>
#include <unordered_map>
using namespace std;

// Use hash map:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 16 ms, faster than 60.15%
// Memory Usage: 9.6 MB, less than 55.22%
class Solution
{
public:
    bool isAnagram(string s, string t)
    {
        if (s.size() != t.size())
        {
            return false;
        }
        unordered_map<char, int> map;

        for (auto c : s)
        {
            map[c]++;
        }

        for (auto c : t)
        {
            if (map.find(c) != map.end())
            {
                if (--map[c] < 0)
                {
                    return false;
                }
            }
            else
            {
                return false;
            }
        }

        return true;
    }
};

// Use hash table:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 12 ms, faster than 83.41%
// Memory Usage: 9.4 MB, less than 82.09%
class Solution2
{
public:
    bool isAnagram(string s, string t)
    {
        if (s.size() != t.size())
        {
            return false;
        }
        int count[26] = {0};

        for (char c : s)
        {
            count[c - 'a']++;
        }

        for (char c : t)
        {
            if (--count[c - 'a'] < 0)
            {
                return false;
            }
        }

        return true;
    }
};