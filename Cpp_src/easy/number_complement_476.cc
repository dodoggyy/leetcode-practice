#include <iostream>
using namespace std;

// Use bit manipulation with iteration
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 6.2 MB, less than 100.00%
class Solution
{
public:
    int findComplement(int num)
    {
        unsigned int mask = ~0;
        while (mask & num)
        {
            mask <<= 1;
        }

        return num ^ ~mask;
    }
};