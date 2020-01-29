#include <iostream>
using namespace std;

// Use iterative:
// Time Complexity: O(logn)
// Space Complexity: O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 8.1 MB, less than 100.00%
class Solution
{
public:
    string toHex(int num)
    {
        char hex[] = {'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'};
        unsigned int n = num;
        string result = "";
        if (n == 0)
        {
            return "0";
        }
        while (n)
        {
            result.insert(result.begin(), hex[(n & 0xf)]);
            n >>= 4;
        }

        return result;
    }
};