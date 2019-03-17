/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
struct TreeNode* insertIntoBST(struct TreeNode* root, int val) {
    
    if(!root) return NULL;
    struct TreeNode *Index = root;
    while(Index)
    {
        if(val > Index->val)
        {
            if(!Index->right) //found the place to insert val
            {
                Index->right = (struct TreeNode *)calloc(1, sizeof(struct TreeNode));
                Index->right->val = val;
                break;
            }
            else
            {
                Index = Index->right;
            }
        }
        else if(val < Index->val)
        {
            if(!Index->left) //found the place to insert val
            {
                Index->left = (struct TreeNode *)calloc(1, sizeof(struct TreeNode));
                Index->left->val = val;
                break;
            }
            else
            {
                Index = Index->left;
            }
        }
    }
    
    return root;
}