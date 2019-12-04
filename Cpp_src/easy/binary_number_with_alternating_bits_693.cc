#include <iostream>
using namespace std;

// Use bit manipulation
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 7.9 MB, less than 100.00%
class Solution
{
public:
    bool hasAlternatingBits(int n)
    {
        long tmp = n ^ (n >> 1);
        return (tmp & (tmp + 1)) == 0;
    }
};