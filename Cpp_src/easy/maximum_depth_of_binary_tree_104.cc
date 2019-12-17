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
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 19.5 MB, less than 63.74%
class Solution
{
public:
    int maxDepth(TreeNode *root)
    {
        if (root == NULL)
        {
            return 0;
        }

        return max(maxDepth(root->left), maxDepth(root->right)) + 1;
    }
};

// Use BFS
// Time Complexity: O(n)
// Space Complexity: O(1)
// Runtime: 12 ms, faster than 58.22%
// Memory Usage: 19.2 MB, less than 92.31%
class Solution2
{
public:
    int maxDepth(TreeNode *root)
    {
        if (root == NULL)
        {
            return 0;
        }

        int result = 0;
        queue<TreeNode *> node_queue;
        node_queue.push(root);
        while (!node_queue.empty())
        {
            int size = node_queue.size();
            for (int i = 0; i < size; i++)
            {
                TreeNode *node = node_queue.front();
                node_queue.pop();
                if (node->left)
                {
                    node_queue.push(node->left);
                }
                if (node->right)
                {
                    node_queue.push(node->right);
                }
            }
            result++;
        }

        return result;
    }
};