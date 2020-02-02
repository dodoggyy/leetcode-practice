#include <iostream>
using namespace std;

// Use linear search:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 56.33%
// Memory Usage: 8.1 MB, less than 90.48%
class Solution
{
public:
    bool isPerfectSquare(int num)
    {
        for (int i = 1; i <= num / i; i++)
        {
            if (i * i == num)
            {
                return true;
            }
        }
        return false;
    }
};

// Use binary search:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 8.1 MB, less than 85.71%
class Solution2
{
public:
    bool isPerfectSquare(int num)
    {
        long left = 0, right = num;

        while (left <= right)
        {
            long mid = left + (right - left) / 2;
            long tmp = mid * mid;
            if (tmp < num)
            {
                left = mid + 1;
            }
            else if (tmp > num)
            {
                right = mid - 1;
            }
            else
            {
                return true;
            }
        }

        return false;
    }
};

// Use arithmetic sequence concept:
// Time Complexity: O(logn)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 8.1 MB, less than 90.48%
class Solution3
{
public:
    bool isPerfectSquare(int num)
    {
        int sub_num = 1;
        while (num > 0)
        {
            num -= sub_num;
            sub_num += 2;
        }

        return (num == 0);
    }
};