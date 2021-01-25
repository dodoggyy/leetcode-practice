#include <iostream>
#include <vector>
using namespace std;

// Use bit manipulation
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 52 ms, faster than 94.35%
// Memory Usage: 9.5 MB, less than 82.93%
class Solution
{
public:
    vector<int> countBits(int num)
    {
        vector<int> result(num + 1, 0);
        for (int i = 1; i <= num; i++)
        {
            result[i] = result[i & (i - 1)] + 1;
        }

        return result;
    }
};