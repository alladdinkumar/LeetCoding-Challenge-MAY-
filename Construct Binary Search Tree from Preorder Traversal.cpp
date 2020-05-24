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
    TreeNode* insert(TreeNode* head, int data)
    {
        if(head==NULL)
        {
            return new TreeNode(data);
        }
        if(head->val > data)
        {
            head->left=insert(head->left,data);
        }
        else
        {
            head->right=insert(head->right,data);
        }
        return head;
    }
    TreeNode* bstFromPreorder(vector<int>& preorder) {
        TreeNode* root=NULL;
        for(auto x: preorder)
        {
            root=insert(root,x);
        }
         return root;   
    }
};
