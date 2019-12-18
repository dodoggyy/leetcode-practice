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

// Use Recursive:
// Time Complexity: O(n)
// Space Complexity: O(h) = O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 9.2 MB, less than 94.55%
class Solution
{
public:
    TreeNode *invertTree(TreeNode *root)
    {
        if (!root)
        {
            return root;
        }
        TreeNode *tmp = root->left;
        root->left = invertTree(root->right);
        root->right = invertTree(tmp);

        return root;
    }
};

// Use Iterative:
// Time Complexity: O(n)
// Space Complexity: O(h) = O(n)
// Runtime: 4 ms, faster than 57.58%
// Memory Usage: 9.1 MB, less than 100.00%
class Solution2
{
public:
    TreeNode *invertTree(TreeNode *root)
    {
        if (!root)
        {
            return root;
        }
        queue<TreeNode *> tree_queue;
        tree_queue.push(root);
        while (!tree_queue.empty())
        {
            TreeNode *node = tree_queue.front();
            tree_queue.pop();
            TreeNode *tmp = node->left;
            node->left = node->right;
            node->right = tmp;
            if (node->left)
            {
                tree_queue.push(node->left);
            }
            if (node->right)
            {
                tree_queue.push(node->right);
            }
        }
        return root;
    }
};