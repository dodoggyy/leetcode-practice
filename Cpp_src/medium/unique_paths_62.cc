#include <iostream>
#include <vector>
using namespace std;

// Use Permutation formula:
// Time Complexity: O(1)
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 5.9 MB, less than 100.00%
class Solution
{
public:
    int uniquePaths(int m, int n)
    {
        long answer = 1;
        int S = m + n - 2;
        int D = n - 1;

        for (int i = 1; i <= D; i++)
        {
            answer = answer * (S - D + i) / i;
        }

        return answer;
    }
};

// Use DP:
// uniquePaths(m,n) = dp[n - 1][m] + dp[n][m-1]
// Time Complexity: O(m*n)
// Space Complexity:O(m*n)
// Runtime: 4 ms, faster than 71.92%
// Memory Usage: 8.2 MB, less than 100.00%
class Solution2
{
public:
    int uniquePaths(int m, int n)
    {
        vector<vector<int>> dp(m, vector<int>(n, 1));
        for (int i = 1; i < m; i++)
        {
            for (int j = 1; j < n; j++)
            {
                dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
            }
        }

        return dp[m - 1][n - 1];
    }
};