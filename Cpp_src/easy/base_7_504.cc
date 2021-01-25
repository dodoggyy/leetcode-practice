#include <iostream>
using namespace std;

// Use recursive:
// Time Complexity: O(logn)
// Space Complexity: O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 8.3 MB, less than 50.00%
class Solution
{
public:
    string convertToBase7(int num)
    {
        if (num < 0)
        {
            return "-" + convertToBase7(-num);
        }
        if (num < 7)
        {
            return to_string(num);
        }
        return convertToBase7(num / 7) + to_string(num % 7);
    }
};

// Use iterative:
// Time Complexity: O(logn)
// Space Complexity: O(1)
// Runtime: 4 ms, faster than 55.63%
// Memory Usage: 8.2 MB, less than 100.00%
class Solution2
{
public:
    string convertToBase7(int num)
    {
        if (num == 0)
        {
            return "0";
        }
        string result = "";
        bool bIsNegative = num < 0;
        num = abs(num);

        while (num)
        {
            result.insert(result.begin(), (char)(num % 7 + '0'));
            num /= 7;
        }

        if (bIsNegative)
        {
            result.insert(result.begin(), (char)'-');
        }

        return result;
    }
};