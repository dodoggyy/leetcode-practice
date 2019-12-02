#include <iostream>
using namespace std;

// Use bit manipulation
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 8.1 MB, less than 95.45%
class Solution
{
public:
    int hammingDistance(int x, int y)
    {
        int tmp = x ^ y;
        int result = 0;
        while (tmp != 0)
        {
            tmp &= tmp - 1;
            result++;
        }
        return result;
    }
};