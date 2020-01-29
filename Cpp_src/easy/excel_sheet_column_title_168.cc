#include <iostream>
using namespace std;

// Use recursive:
// Time Complexity: O(logn)
// Space Complexity: O(logn)
// Runtime: 4 ms, faster than 48.81%
// Memory Usage: 8.1 MB, less than 100.00%
class Solution
{
public:
    string convertToTitle(int n)
    {
        if (n == 0)
        {
            return "";
        }
        // symbol start from 1 rather than 0
        // so that need to minus 1
        n--;
        return convertToTitle(n / 26) + (char)(n % 26 + 'A');
    }
};

// Use iterative:
// Time Complexity: O(logn)
// Space Complexity: O(1)
// Runtime: 4 ms, faster than 48.81%
// Memory Usage: 8.2 MB, less than 72.73%
class Solution2
{
public:
    string convertToTitle(int n)
    {
        string result;

        while (n)
        {
            n--;
            result = (char)(n % 26 + 'A') + result;
            n /= 26;
        }

        return result;
    }
};