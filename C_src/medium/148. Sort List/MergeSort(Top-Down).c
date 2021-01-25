/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
typedef struct ListNode ListNode_t;

ListNode_t *FindMidFromList(ListNode_t *head, ListNode_t *end)
{
    if(head == NULL || end == NULL)
    {
        return NULL;
    }
    if(head == end) // only one element in the list
    {
        return head;
    }
    // find mid id from list
    ListNode_t *mid = head;
    int s32Idx = 0;
    //total lenth of list
    while(mid != end)
    {
        mid = mid->next;
        s32Idx++;
    }
    //mid idx = total length divide by 2
    s32Idx /=2;

    mid = head;
    while(s32Idx)
    {
        mid = mid->next;
        s32Idx--;
    }
    return mid;
}
void MoveTheNode(ListNode_t **DestNode, ListNode_t **SourceNode)
{
    ListNode_t *TempNode = *SourceNode;
    *SourceNode = (*SourceNode)->next; // Source list point to next node
    TempNode->next = *DestNode;
    *DestNode = TempNode;
    return;
}

ListNode_t *MergeList(ListNode_t *LeftList, ListNode_t *RightList)
{

    ListNode_t DummyNode; // use dummy node to combine two list.
    memset(&DummyNode, 0, sizeof(ListNode_t));
    ListNode_t *head = &DummyNode;
    while(1)
    {
        if(LeftList == NULL)
        {
            head->next = RightList;
            break;
        }
        else if(RightList == NULL)
        {
            head->next = LeftList;
            break;
        }

        if(LeftList->val > RightList->val)
        {
            MoveTheNode(&(head->next), &RightList);
        }
        else
        {
            MoveTheNode(&(head->next), &LeftList);
        }
        head = head->next;
    }
    head = DummyNode.next; // ignore dummy node.
    return head;


}



struct ListNode* sortList(struct ListNode* head) {
    
    if(head == NULL)
        return NULL;

    if(head->next == NULL) // only one node in the list.
        return head;

    ListNode_t *end = head;
    while(end->next != NULL)
    {
        end = end->next;
    }

    ListNode_t *mid = FindMidFromList(head, end);
    ListNode_t *midnextnode = mid->next;
    //split list
    mid->next = NULL;
    return MergeList(sortList(head), sortList(midnextnode));

}