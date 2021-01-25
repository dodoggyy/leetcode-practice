/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
struct TreeNode* searchBST(struct TreeNode* root, int val) {
    
    if(!root) return NULL;
    
    while(root)
    {
        if(val == root->val) return root;
        else if(val > root->val) root = root->right;
        else //less than root val
            root = root->left;
    }
    //can't find target in given binary search tree.
    return NULL;
}