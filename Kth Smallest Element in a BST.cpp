/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    int kthsmall(TreeNode* root, int &count , int k)
    {
        if (root==NULL)
            return -1;
        int left= kthsmall(root->left,count,k);
        if(left!=-1)
            return left;
        count++;
        if(count==k)
        {
            return root->val;
        }
        return kthsmall(root->right,count,k);
    }
    int kthSmallest(TreeNode* root, int k) {
        int count=0;
        return kthsmall(root,count,k);
        
    }
};
