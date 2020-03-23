#include <iostream>
#include <vector>
using namespace std;

// Use DP:
// uniquePaths(m,n) = dp[n - 1][m] + dp[n][m-1]
// Time Complexity: O(m*n)
// Space Complexity:O(m*n)
// Runtime: 4 ms, faster than 75.19%
// Memory Usage: 7.1 MB, less than 100.00%
class Solution
{
public:
    int uniquePathsWithObstacles(vector<vector<int>> &obstacleGrid)
    {
        int m = obstacleGrid.size(), n = obstacleGrid[0].size();
        vector<vector<long>> dp(m, vector<long>(n, 0)); // error when type is int for some case

        for (int i = 0; i < m; i++)
        {
            for (int j = 0; j < n; j++)
            {
                if (obstacleGrid[i][j] == 1)
                {
                    dp[i][j] = 0;
                }
                else
                {
                    if (i == 0 && j == 0)
                    {
                        dp[i][j] = 1;
                    }
                    else if (i == 0)
                    {
                        dp[i][j] = dp[i][j - 1];
                    }
                    else if (j == 0)
                    {
                        dp[i][j] = dp[i - 1][j];
                    }
                    else
                    {
                        dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
                    }
                }
            }
        }

        return dp[m - 1][n - 1];
    }
};

// optimize DP:
// uniquePaths(m,n) = dp[n - 1][m] + dp[n][m-1]
// Time Complexity: O(m*n)
// Space Complexity:O(m*n)
// Runtime: 4 ms, faster than 75.19%
// Memory Usage: 7.1 MB, less than 100.00%
class Solution2
{
public:
    int uniquePathsWithObstacles(vector<vector<int>> &obstacleGrid)
    {
        int m = obstacleGrid.size(), n = obstacleGrid[0].size();
        vector<vector<long>> dp(m + 1, vector<long>(n + 1, 0)); // error when type is int for some case
        dp[0][1] = 1;                                           // equal dp[1][0] = 1
        for (int i = 1; i <= m; i++)
        {
            for (int j = 1; j <= n; j++)
            {
                if (!obstacleGrid[i - 1][j - 1])
                {
                    dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
                }
            }
        }

        return dp[m][n];
    }
};

// optimize DP 2:
// Time Complexity: O(m*n)
// Space Complexity:O(n)
// Runtime: 4 ms, faster than 75.19%
// Memory Usage: 7 MB, less than 100.00%
class Solution3
{
public:
    int uniquePathsWithObstacles(vector<vector<int>> &obstacleGrid)
    {
        int n = obstacleGrid[0].size();
        vector<long> dp(n, 0);
        dp[0] = 1;

        for (vector<int> row : obstacleGrid)
        {
            for (int j = 0; j < n; j++)
            {
                if (row[j] == 1)
                {
                    dp[j] = 0;
                }
                else if (j == 0)
                {
                    continue;
                }
                else
                {
                    dp[j] += dp[j - 1];
                }
            }
        }

        return dp[n - 1];
    }
};