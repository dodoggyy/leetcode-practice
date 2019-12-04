#include <iostream>
using namespace std;

// Use bit manipulation
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 56.73%
// Memory Usage: 8.1 MB, less than 80.95%
class Solution
{
public:
    bool isPowerOfFour(int num)
    {
        return (num > 0) && ((num & (num - 1)) == 0) && (num & 0x55555555) != 0;
    }
};