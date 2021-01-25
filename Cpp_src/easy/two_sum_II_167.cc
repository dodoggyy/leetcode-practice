#include <iostream>
#include <vector>
using namespace std;

// Use dual index
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 4 ms, faster than 96.86%
// Memory Usage: 9.5 MB, less than 80.39%
class Solution
{
public:
    vector<int> twoSum(vector<int> &numbers, int target)
    {
        vector<int> result;
        int index_left = 0, index_right = numbers.size() - 1;
        while (index_left < index_right)
        {
            int tmp = numbers.at(index_left) + numbers.at(index_right);
            if (tmp == target)
            {
                result.push_back(index_left + 1);
                result.push_back(index_right + 1);
                return result;
            }
            else if (tmp < target)
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