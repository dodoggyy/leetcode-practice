#include <iostream>
using namespace std;

// Use bit manipulation
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 8.1 MB, less than 91.67%
class Solution
{
public:
    uint32_t reverseBits(uint32_t n)
    {
        uint32_t result = 0;
        for (int i = 0; i < 32; i++)
        {
            result <<= 1;
            result |= (1 & n);
            n >>= 1;
        }

        return result;
    }
};