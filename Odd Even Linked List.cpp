/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* oddEvenList(ListNode* head) {
        
        if(head==NULL || head->next==NULL)
            return head;
        ListNode* oddend=head;
        ListNode* evenstart=head->next;
        ListNode* tmp=head;
        ListNode* crawlp=evenstart;
        while(crawlp!=NULL && crawlp->next!=NULL )
        {
            ListNode* crawl=crawlp->next;
            crawlp->next=crawl->next;
            oddend->next=crawl;
            crawl->next=evenstart;
            oddend=oddend->next;
            crawlp=crawlp->next;
            
        }
        head=tmp;
        return head;
        
    }
};
