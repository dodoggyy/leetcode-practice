#include <iostream>
#include <vector>
using namespace std;

// Use two pointers
// Time Complexity: O(n)
// Space Complexity: O(1)
// Runtime: 64 ms, faster than 37.01%
// Memory Usage: 13 MB, less than 100.00%
class Solution
{
public:
    void reverseString(vector<char> &s)
    {
        for (int i = 0; i < (s.size() / 2); i++)
        {
            swap(s, i, s.size() - i - 1);
        }
    }

    void swap(vector<char> &s, int i, int j)
    {
        char tmp = s[i];
        s[i] = s[j];
        s[j] = tmp;
    }
};