#include <iostream>
#include <vector>
using namespace std;

// Use search:
// Time Complexity: O(m+n)
// Space Complexity:O(1)
// Runtime: 56 ms, faster than 98.29%
// Memory Usage: 12.9 MB, less than 73.33%
class Solution
{
public:
    bool searchMatrix(vector<vector<int>> &matrix, int target)
    {
        if (matrix.size() == 0)
        {
            return false;
        }
        int size_row = matrix.size(), size_col = matrix[0].size();
        int row = 0, col = size_col - 1;

        while (row < size_row && col >= 0)
        {
            if (matrix[row][col] > target)
            {
                col--;
            }
            else if (matrix[row][col] < target)
            {
                row++;
            }
            else
            {
                return true;
            }
        }

        return false;
    }
};