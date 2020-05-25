/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head==nil || head.Next==nil{
        return head
    }
    oddend:=head
    evenstart:=head.Next
    tmp:=head
    crawlp:=evenstart
    for ;crawlp!=nil && crawlp.Next!= nil;{
        crawl:=crawlp.Next
        crawlp.Next=crawl.Next
        oddend.Next=crawl
        crawl.Next=evenstart
        oddend=oddend.Next
        crawlp=crawlp.Next
    }
    head=tmp
    return head
    
}