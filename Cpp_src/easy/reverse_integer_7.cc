#include <iostream>
using namespace std;

// Use Pop and Push Digits & Check before Overflow
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 69.15%
// Memory Usage: 8.2 MB, less than 93.33%
class Solution
{
public:
    int reverse(int x)
    {
        if (x == INT_MIN)
        {
            return 0;
        }
        int result = 0;

        while (x != 0)
        {
            if ((result > (INT_MAX / 10)) || (result < (INT_MIN / 10)))
            {
                return 0;
            }

            result = result * 10 + x % 10;
            x /= 10;
        }

        return result;
    }
};