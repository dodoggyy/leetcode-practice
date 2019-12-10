#include <stdio.h>
#include <stdlib.h>

// Definition for singly-linked list.
struct ListNode
{
    int val;
    struct ListNode *next;
};

// Use iterative:
// Time Complexity: O(MIN(m,n))
// Space Complexity:O(m+n)
// Runtime: 4 ms, faster than 78.66%
// Memory Usage: 7.4 MB, less than 45.45%
struct ListNode *mergeTwoLists(struct ListNode *l1, struct ListNode *l2)
{
    struct ListNode stResult = {.val = 0};
    struct ListNode *pstCur = &stResult;

    while (l1 != NULL && l2 != NULL)
    {
        if (l1->val < l2->val)
        {
            pstCur->next = l1;
            l1 = l1->next;
        }
        else
        {
            pstCur->next = l2;
            l2 = l2->next;
        }
        pstCur = pstCur->next;
    }
    if (l1 != NULL)
    {
        pstCur->next = l1;
    }
    if (l2 != NULL)
    {
        pstCur->next = l2;
    }
    return stResult.next;
}

// Use recursive:
// Time Complexity: O(MIN(m,n))
// Space Complexity:O(1)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 7.5 MB, less than 45.45%
struct ListNode *mergeTwoLists2(struct ListNode *l1, struct ListNode *l2)
{
    if (l1 == NULL)
    {
        return l2;
    }
    if (l2 == NULL)
    {
        return l1;
    }

    if (l1->val < l2->val)
    {
        l1->next = mergeTwoLists2(l1->next, l2);
        return l1;
    }
    else
    {
        l2->next = mergeTwoLists2(l1, l2->next);
        return l2;
    }
}