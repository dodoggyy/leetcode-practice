#include <iostream>
#include <queue>
using namespace std;

// Definition for a binary tree node.
struct TreeNode
{
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

// Use DFS
// Time Complexity: O(n)
// Space Complexity: O(1)
// Runtime: 12 ms, faster than 84.47%
// Memory Usage: 17.1 MB, less than 96.61%
class Solution
{
public:
    bool isBalanced(TreeNode *root)
    {
        depth(root);
        return result;
    }

private:
    bool result = true;
    int depth(TreeNode *root)
    {
        if (!root)
        {
            return 0;
        }
        int left = depth(root->left);
        int right = depth(root->right);

        if (abs(left - right) > 1)
        {
            result = false;
        }

        return max(left, right) + 1;
    }
};