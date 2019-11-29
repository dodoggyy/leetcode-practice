#include <iostream>
#include <unordered_map>
#include <unordered_set>
using namespace std;

// Use slide window with hash set
// Time Complexity: O(2n) = O(n)
// Space Complexity:O(k)
// Runtime: 32 ms, faster than 33.86%
// Memory Usage: 13.3 MB, less than 30.85%
class Solution
{
public:
    int lengthOfLongestSubstring(string s)
    {
        unordered_set<char> hash_set;
        unordered_set<char>::iterator iter;
        int result = 0;
        int i = 0, j = 0, length = s.size();
        while (i < length && j < length)
        {
            iter = hash_set.find(s.at(j));
            if (iter == hash_set.end())
            {
                hash_set.insert(s.at(j++));
                result = max(result, j - i);
            }
            else
            {
                hash_set.erase(s.at(i++));
            }
        }

        return result;
    }
};

// Use slide window (Assuming ASCII 128)
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 12 ms, faster than 83.19%
// Memory Usage: 9 MB, less than 98.51%
class Solution2
{
public:
    int lengthOfLongestSubstring(string s)
    {
        int result = 0;
        int length = s.size();
        int index[128] = {0};
        for (int i = 0, j = 0; j < length; j++)
        {
            i = max(index[(int)s.at(j)], i);
            result = max(result, j - i + 1);
            index[(int)s.at(j)] = j + 1; // 1 for after same char position
        }
        return result;
    }
};