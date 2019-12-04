#include <iostream>
using namespace std;

// Use bit manipulation
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 56.82%
// Memory Usage: 8.1 MB, less than 68.00%
class Solution
{
public:
    bool isPowerOfTwo(int n)
    {
        return ((n > 0) && ((n & (n - 1)) == 0));
    }
};