#include <iostream>
#include <unordered_map>
#include <vector>
using namespace std;

// Use hash map
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 8 ms, faster than 92.92%
// Memory Usage: 10.1 MB, less than 54.77%
class Solution
{
public:
    vector<int> twoSum(vector<int> &nums, int target)
    {
        vector<int> result;
        unordered_map<int, int> hash_map;
        unordered_map<int, int>::iterator iter;
        for (int i = 0; i < nums.size(); i++)
        {
            iter = hash_map.find(target - nums[i]);
            if (iter != hash_map.end())
            {
                result.push_back(iter->second);
                result.push_back(i);
                return result;
            }
            else
            {
                hash_map.insert(pair<int, int>(nums[i], i));
            }
        }
        return result;
    }
};