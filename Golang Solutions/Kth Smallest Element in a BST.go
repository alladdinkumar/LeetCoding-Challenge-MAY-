/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthsmall(root *TreeNode,count *int, k int) int {
    if root==nil{
        return -1       
    }   
    var left int
    left= kthsmall(root.Left,count,k)
    if left!=-1{
        return left
    }
    *count++        
    if *count==k{
        return root.Val
    }
    return kthsmall(root.Right,count,k);
}
func kthSmallest(root *TreeNode ,  k int)int {
    count:=0
    return kthsmall(root,&count,k)
}