#include <stdio.h>
#include <stdlib.h>

// Use Two Pointers:
// Time Complexity: O(m+n)
// Space Complexity:O(1)
// Runtime: 36 ms, faster than 75.00%
// Memory Usage: 15.3 MB, less than 33.33%

// Definition for singly-linked list.
struct ListNode
{
    int val;
    struct ListNode *next;
};

struct ListNode *getIntersectionNode(struct ListNode *headA, struct ListNode *headB)
{
    struct ListNode *stNodeA = headA, *stNodeB = headB;
    while (stNodeA != stNodeB)
    {
        stNodeA = stNodeA ? stNodeA->next : headB;
        stNodeB = stNodeB ? stNodeB->next : headA;
    }
    return stNodeA;
}