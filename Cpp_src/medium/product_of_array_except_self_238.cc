#include <iostream>
#include <vector>
using namespace std;

// Use dual array record:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 36 ms, faster than 98.25%
// Memory Usage: 13 MB, less than 21.21%
class Solution
{
public:
    vector<int> productExceptSelf(vector<int> &nums)
    {
        int size = nums.size();
        vector<int> result(size, 0);
        int L[size], R[size];

        L[0] = 1, R[size - 1] = 1;
        for (int i = 1; i < size; i++)
        {
            L[i] = L[i - 1] * nums[i - 1];
            R[size - i - 1] = R[size - i] * nums[size - i];
        }

        for (int i = 0; i < size; i++)
        {
            result[i] = L[i] * R[i];
        }

        return result;
    }
};

// Use 1 parameter and result array record:
// Time Complexity: O(n)
// Space Complexity:O(1)
// Runtime: 40 ms, faster than 85.61%
// Memory Usage: 12.5 MB, less than 77.27%
class Solution2
{
public:
    vector<int> productExceptSelf(vector<int> &nums)
    {
        int size = nums.size();
        vector<int> result(size, 1);

        for (int i = 1; i < size; i++)
        {
            result[i] = result[i - 1] * nums[i - 1];
        }

        int R = 1;
        for (int i = size - 1; i >= 0; i--)
        {
            result[i] *= R;
            R *= nums[i];
        }

        return result;
    }
};