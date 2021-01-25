/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2) {
    struct ListNode *pList1 = l1;
    struct ListNode *pList2 = l2;
    struct ListNode *pSumList = NULL;
    struct ListNode *pHeadList = (struct ListNode *)malloc(sizeof(struct ListNode));
    memset(pHeadList, 0 , sizeof(struct ListNode));
    pSumList = pHeadList;
    int Sum = 0;
    int Carrybit = 0;

    while(pList1 || pList2)
    {
        int Val_1 = pList1 ? pList1->val : 0;
        int Val_2 = pList2 ? pList2->val : 0;
        Sum = Val_1 + Val_2 + Carrybit;
        Carrybit = Sum > 9 ? 1 : 0;
        pSumList->val = Sum % 10;
        pList1 = pList1 ? pList1->next : NULL;
        pList2 = pList2 ? pList2->next : NULL;
        if(pList1 || pList2)
        {
            struct ListNode *pTempList = (struct ListNode *)malloc(sizeof(struct ListNode));
            memset(pTempList, 0 , sizeof(struct ListNode));
            pSumList->next = pTempList;
            pSumList = pTempList;

        }
    }
    if(Carrybit)
    {
        struct ListNode *pTempList = (struct ListNode *)malloc(sizeof(struct ListNode));
        memset(pTempList, 0 , sizeof(struct ListNode));
        pTempList->val = Carrybit;
        pSumList->next = pTempList;
        pSumList = pTempList;
    }
    return pHeadList;
}