/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func insert(head *TreeNode, data int) *TreeNode{
        if head==nil{
            var n *TreeNode=new(TreeNode)
            n.Val=data
            return n
        }
        if head.Val > data{
            head.Left=insert(head.Left,data);
        }else{
            head.Right=insert(head.Right,data);
        }
        return head
    }
func bstFromPreorder(preorder []int) *TreeNode {
        var root *TreeNode
    for _,x:=range preorder{
            root=insert(root,x)
        }
    return root;  
}