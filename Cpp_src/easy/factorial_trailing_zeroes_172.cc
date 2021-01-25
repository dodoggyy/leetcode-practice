#include <iostream>
#include <vector>
using namespace std;

// Use iterative:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 8.2 MB, less than 85.71%
class Solution
{
public:
    int trailingZeroes(int n)
    {
        int result = 0;
        while (n)
        {
            n /= 5;
            result += n;
        }
        return result;
    }
};

// Use recursive:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 8.3 MB, less than 78.57%
class Solution2
{
public:
    int trailingZeroes(int n)
    {
        return (n == 0) ? 0 : (n / 5 + trailingZeroes(n / 5));
    }
};