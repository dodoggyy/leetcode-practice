#include <iostream>
using namespace std;

// Definition for singly-linked list.
struct ListNode
{
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

// Use iterative
// Time Complexity: O(MAX(m,n))
// Space Complexity:O(m+n)
// Runtime: 8 ms, faster than 78.40%
// Memory Usage: 8.9 MB, less than 72.13%
class Solution
{
public:
    ListNode *mergeTwoLists(ListNode *l1, ListNode *l2)
    {
        ListNode *result = new ListNode(0);
        ListNode *cur = result;
        while (l1 != NULL && l2 != NULL)
        {
            if (l1->val > l2->val)
            {
                cur->next = l2;
                l2 = l2->next;
            }
            else
            {
                cur->next = l1;
                l1 = l1->next;
            }
            cur = cur->next;
        }
        if (l1 != NULL)
        {
            cur->next = l1;
        }
        if (l2 != NULL)
        {
            cur->next = l2;
        }

        return result->next;
    }
};

// Use recursive
// Time Complexity: O(MAX(m,n))
// Space Complexity:O(1)
// Runtime: 8 ms, faster than 78.40%
// Memory Usage: 8.9 MB, less than 77.05%
class Solution2
{
public:
    ListNode *mergeTwoLists(ListNode *l1, ListNode *l2)
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
            l1->next = mergeTwoLists(l1->next, l2);
            return l1;
        }
        else
        {
            l2->next = mergeTwoLists(l1, l2->next);
            return l2;
        }
    }
};