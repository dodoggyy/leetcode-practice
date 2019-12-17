#include <stdio.h>
#include <stdlib.h>

// Definition for a binary tree node.
struct TreeNode
{
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

// Use DFS:
// Time Complexity: O(n)
// Space Complexity:O(logn)
// Runtime: 12 ms, faster than 10.55%
// Memory Usage: 9.1 MB, less than 100.00%
#define MAX(a, b) ((a) > (b) ? (a) : (b))

int maxDepth(struct TreeNode *root)
{
    if (!root)
    {
        return 0;
    }

    int left = maxDepth(root->left);
    int right = maxDepth(root->right);

    return MAX(left, right) + 1;
}