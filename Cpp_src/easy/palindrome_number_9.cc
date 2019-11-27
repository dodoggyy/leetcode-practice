#include <iostream>
using namespace std;

// Use loop reverse
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 98.06%
// Memory Usage: 8.1 MB, less than 88.18%
class Solution
{
public:
    bool isPalindrome(int x)
    {
        if (x < 0)
        {
            return false;
        }
        int tmp = 0, origin = x;
        while (origin != 0)
        {
            if (tmp > INT_MAX / 10)
            {
                return false;
            }
            tmp = tmp * 10 + (origin % 10);
            origin /= 10;
        }

        return tmp == x;
    }
};