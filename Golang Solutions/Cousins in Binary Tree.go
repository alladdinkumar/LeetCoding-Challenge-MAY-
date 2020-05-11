/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
    q:=[]*TreeNode{}
    q_val:=[]int{}
    q=append(q,root)
    q_val=append(q_val,0)
    var depth_x,depth_y int
    cnt,d:=0,0
    for ;len(q)!=0 && cnt<=2; {
        var n *TreeNode= new(TreeNode)
        n=q[0]
        d=q_val[0]
        q=q[1:]
        q_val=q_val[1:]
        if(n.Left != nil && n.Right != nil){
            if((n.Left.Val==x && n.Right.Val==y) ||(n.Left.Val==y && n.Right.Val==x)){
                return false;
            }         
        }
        if n.Val==x{
            depth_x=d
            cnt++
        }else if(n.Val == y){
            depth_y=d
            cnt++
        }
        if(n.Left != nil){
            q=append(q,n.Left)
            q_val=append(q_val,d+1)  
        }
        if(n.Right != nil){
             q=append(q,n.Right)
            q_val=append(q_val,d+1) 
        }
        
    }
    if depth_x==depth_y{
        return true
    }
    return false      
}