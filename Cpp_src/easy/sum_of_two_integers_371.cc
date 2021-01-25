#include <iostream>
using namespace std;

// Use bit manipulation with iteration
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 50.43%
// Memory Usage: 8.1 MB, less than 100.00%
class Solution
{
public:
    int getSum(int a, int b)
    {
        while (b)
        {
            int carry = ((a & b) & 0xffffffff) << 1;
            a ^= b;
            b = carry;
        }
        return a;
    }
};

// Use bit manipulation with recursive
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 50.43%
// Memory Usage: 8.2 MB, less than 58.33%
class Solution2
{
public:
    int getSum(int a, int b)
    {
        return (b == 0) ? a : getSum((a ^ b), (((a & b) & (0xffffffff)) << 1));
    }
};