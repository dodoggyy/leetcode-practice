#include <iostream>
#include <vector>
using namespace std;

// Use two pointers
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 16 ms, faster than 95.98%
// Memory Usage: 9.9 MB, less than 29.90%
class Solution
{
public:
    int maxArea(vector<int> &height)
    {
        int result = 0;
        int index_left = 0, index_right = height.size() - 1;
        while (index_left < index_right)
        {
            result = max(result, min(height[index_left], height[index_right]) * (index_right - index_left));
            if (height[index_left] < height[index_right])
            {
                index_left++;
            }
            else
            {
                index_right--;
            }
        }

        return result;
    }
};