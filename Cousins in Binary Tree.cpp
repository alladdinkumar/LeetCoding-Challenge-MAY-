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
    
    bool isCousins(TreeNode* root, int x, int y) {
        queue<pair<TreeNode*,int>>q;
        pair<TreeNode*,int>k,l;
        k.first=root;
        k.second=0;
        q.push(k);
        int depth_x,depth_y;
        int count=0;
        
        while(!q.empty() && count!=2)
        {
            k=q.front();
            q.pop();
            TreeNode* n=k.first;
            int depth=k.second;
            if(n->left!= NULL && n->right!= NULL)
            {
                if((n->left->val==x && n->right->val==y) ||(n->left->val==y && n->right->val==x))
                    return false;
            }
            if(n->val==x)
            {
                depth_x=depth;
                count++;
            }
            else if(n->val==y)
            {
                depth_y=depth;
                count++;
            }
            if(n->left != NULL)
            {
                k.first=n->left;
                k.second=depth+1;
                q.push(k);
                
            }
            if(n->right != NULL)
            {
                k.first=n->right;
                k.second=depth+1;
                q.push(k);
            }
        }
        
        return depth_x==depth_y?true:false;
            
    }
};
