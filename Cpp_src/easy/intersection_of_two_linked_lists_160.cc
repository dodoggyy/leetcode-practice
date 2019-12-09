#include <iostream>
using namespace std;

// Use Two Pointers
// Time Complexity: O(m+n)
// Space Complexity:O(1)
// Runtime: 44 ms, faster than 95.69%
// Memory Usage: 16.7 MB, less than 90.74%

// Definition for singly-linked list.
struct ListNode
{
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution
{
public:
    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB)
    {
        if (headA == NULL || headB == NULL)
        {
            return NULL;
        }

        ListNode *tmp_a = headA;
        ListNode *tmp_b = headB;

        while (tmp_a != tmp_b)
        {
            tmp_a = (tmp_a == NULL) ? headB : tmp_a->next;
            tmp_b = (tmp_b == NULL) ? headA : tmp_b->next;
        }

        return tmp_a;
    }
};